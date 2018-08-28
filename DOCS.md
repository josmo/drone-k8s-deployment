Use this drone plugin to upgrade a k8s deployment 

The following parameters are used to configure this plugin:

- `url` - Url to your cluster server
- `token` - Token used to connect to the cluster
- `cert` - Cert to connect to cluster //Not implemented
- `insecure` - Allow for insecure cluster connection //Not verified
- `deployment_names` - Name(s) of the deployment to update
- `container_names` - Container(s) name in the deployment to update
- `namespaces` - Namespace(s) (will use default if not set)  
- `docker_image` - New image to assign to container in the deployment, including tag (`drone/drone:latest`)
- `date_label` - Label where the date of the deployment took place. Use it when you want to force container upgrading, even if you docker image tag has not changed.


The following is a sample k8s deployment configuration in your `.drone.yml` file:

```yaml
pipeline:
  deploy:
    image: peloton/drone-k8s-deployment
    url: https://k8s.server/
    token: asldkfj
    insecure: false
    deployment_names: mything
    container_names: mything
    namespaces: mynamespace
    docker_image: drone/drone:latest
```

Or with multiples

```yaml
pipeline:
  deploy:
    image: peloton/drone-k8s-deployment
    url: https://k8s.server/
    token: asldkfj
    insecure: false
    deployment_names: [mything, mything2]
    container_names: mything
    namespaces: [mynamespace, anothernamspace]
    docker_image: drone/drone:latest
```

It's not recommended to update multiples of namespaces, deployment_names, and conatiner_names.  Try to keep things simple.

If you want to add secrets for the token or url use KUBERNETES_TOKEN, KUBERNETES_URL

```yaml
pipeline:
  deploy:
    image: peloton/drone-k8s-deployment
    token: asldkfj
    insecure: false
    deployment_names: [mything, mything2]
    container_names: mything
    namespaces: [mynamespace, anothernamspace]
    docker_image: drone/drone:latest
    secrets: [kubernetes_url, kubernetes_token]
```

If you want to force docker container upgrading, there is a hack which allows you to do this even if your image name/tag has not changed. Put to the data_label key ENV variable from Drone CI with a timestamp, such as this:

```yaml
pipeline:
  deploy:
    image: peloton/drone-k8s-deployment
    token: asldkfj
    insecure: false
    deployment_names: [mything, mything2]
    container_names: mything
    namespaces: [mynamespace, anothernamspace]
    docker_image: drone/drone:latest
    date_label: "${DRONE_BUILD_FINISHED}"
    secrets: [kubernetes_url, kubernetes_token]
```

The list of available Drone CI env variables you could find [here](http://docs.drone.io/environment-reference/)