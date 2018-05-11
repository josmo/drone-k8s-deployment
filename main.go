package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var build string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "drone k8s deployment"
	app.Usage = "drone k8s deployemtn"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.0+%s", build)

	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:   "url",
			Usage:  "url to the k8s api",
			EnvVar: "PLUGIN_URL",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "kubernetes token",
			EnvVar: "PLUGIN_TOKEN, KUBERNETES_TOKEN",
		},
		cli.BoolFlag{
			Name:   "insecure",
			Usage:  "Insecure connection",
			EnvVar: "PLUGIN_INSECURE",
		},
		cli.StringSliceFlag{
			Name:   "deployment-names",
			Usage:  "K8s deployment names",
			EnvVar: "PLUGIN_DEPLOYMENT_NAMES",
		},
		cli.StringSliceFlag{
			Name:   "container-names",
			Usage:  "K8s container names for the deployment",
			EnvVar: "PLUGIN_CONTAINER_NAMES",
		},
		cli.StringSliceFlag{
			Name:   "namespaces",
			Usage:  "K8s deployment namspaces",
			EnvVar: "PLUGIN_NAMESPACES",
		},
		cli.StringFlag{
			Name:   "docker-image",
			Usage:  "image to use",
			EnvVar: "PLUGIN_DOCKER_IMAGE",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		URL:             c.String("url"),
		Token:           c.String("token"),
		Insecure:        c.Bool("insecure"),
		DeploymentNames: c.StringSlice("deployment-names"),
		ContainerNames:  c.StringSlice("container-names"),
		NameSpaces:      c.StringSlice("namespaces"),
		DockerImage:     c.String("docker-image"),
	}
	return plugin.Exec()
}
