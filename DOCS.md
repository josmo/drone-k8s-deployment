Use this drone plugin to upgrade a k8s deployment 

The following parameters are used to configure this plugin:

- `url` - url to your cluster server
- `token` - Token used to connect to the cluster
- `cert` - cert to connect to cluster //Not implemented
- `insecure` - allow for insecure cluster connection //Not verified
- `deployment_names` - name(s) of the deployment to update
- `container_names` - container(s) name in the deployment to update
- `namespaces` - namespace(s) (will use default if not set)  
- `docker_image` - new image to assign to container in the deployment, including tag (`drone/drone:latest`)


The following is a sample k8s deployment configuration in your `.drone.yml` file:

```yaml
deploy:
  k8s-deployment:
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
deploy:
  k8s-deployment:
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

if you want to add secrets for the token it's KUBERNETES_TOKEN
