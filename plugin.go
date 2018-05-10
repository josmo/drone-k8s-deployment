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
	URL            string
	Token          string
	Insecure       bool
	DeploymentName string
	ContainerName  string
	NameSpace      string
	DockerImage    string
}

func (p *Plugin) Exec() error {
	log.Info("Drone k8s deployment plugin")

	if p.URL == "" || p.Token == "" || p.DeploymentName == "" || p.ContainerName == "" || p.DockerImage == "" {
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

	namespace := apiv1.NamespaceDefault
	if p.NameSpace != "" {
		namespace = p.NameSpace
	}
	log.Infof("Updating in namespace: %s", namespace)
	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	// More Info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		log.Infof("Updating deployment: %s", p.DeploymentName)
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := deploymentsClient.Get(p.DeploymentName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of deployment: %v", getErr))
		}

		for i := range result.Spec.Template.Spec.Containers {
			// Find the matching conatiner to update
			if result.Spec.Template.Spec.Containers[i].Name == p.ContainerName {
				log.Infof("Updating container: %s with image: %s", p.ContainerName, p.DockerImage)
				result.Spec.Template.Spec.Containers[i].Image = p.DockerImage
			}
		}

		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})
	if retryErr != nil {
		return fmt.Errorf("update failed: %v", retryErr)
	}
	log.Infof("Updated deployment: %s", p.DeploymentName)
	return nil
}
