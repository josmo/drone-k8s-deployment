package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
)

type Plugin struct {
	URL             string
	Token           string
	Insecure        bool
	DeploymentNames []string
	ContainerNames  []string
	NameSpaces      []string
	DockerImage     string
}

func (p *Plugin) Exec() error {
	log.Info("Drone k8s deployment plugin")

	if p.URL == "" || p.Token == "" || len(p.DeploymentNames) <= 0 || len(p.ContainerNames) <= 0 || p.DockerImage == "" {
		return errors.New("eek: only the Namespace, Cert and Ignore flag are not required")
	}

	config := &rest.Config{
		Host:            p.URL,
		BearerToken:     p.Token,
		TLSClientConfig: rest.TLSClientConfig{Insecure: p.Insecure},
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	namespaces := []string{apiv1.NamespaceDefault}
	if len(p.NameSpaces) > 0 {
		namespaces = p.NameSpaces
	}
	log.Infof("Updating container(s): %v in deployment(s): %v in namespace(s): %v", p.ContainerNames, p.DeploymentNames, namespaces)
	for _, namespace := range namespaces {
		log.Infof("Updating in namespace: %s", namespace)
		deploymentsClient := clientset.AppsV1().Deployments(namespace)

		for _, deploymentName := range p.DeploymentNames {
			log.Infof("Updating deployment: %s", deploymentName)
			// More Info:
			// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#concurrency-control-and-consistency
			retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				// Retrieve the latest version of Deployment before attempting update
				// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
				result, getErr := deploymentsClient.Get(deploymentName, metav1.GetOptions{})
				if getErr != nil {
					return fmt.Errorf("failed to get latest version of deployment: %v", getErr)
				}

				for i := range result.Spec.Template.Spec.Containers {
					// Find the matching containers to update
					for _, containerName := range p.ContainerNames {
						if result.Spec.Template.Spec.Containers[i].Name == containerName {
							log.Infof("Updating container: %s with image: %s", containerName, p.DockerImage)
							result.Spec.Template.Spec.Containers[i].Image = p.DockerImage
						}
					}
				}

				_, updateErr := deploymentsClient.Update(result)
				return updateErr
			})
			if retryErr != nil {
				return fmt.Errorf("update failed: %v", retryErr)
			}
			log.Infof("Updated container(s): %v in deployment: %s", p.ContainerNames, deploymentName)
		}
	}
	log.Infof("Updated container(s): %v in deployments: %v in namespaces: %v", p.ContainerNames, p.DeploymentNames, namespaces)
	return nil
}
