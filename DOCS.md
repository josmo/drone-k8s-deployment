Use this drone plugin to upgrade a k8s deployment 

The following parameters are used to configure this plugin:

- `url` - url to your cluster server
- `token` - Token used to connect to the cluster
- `cert` - cert to connect to cluster //Not implemented
- `insecure` - allow for insecure cluster connection //Not verified
- `deployment_name` - name of the deployment to update
- `container_name` - container name in the deployment to update
- `namespace` - namespace (will use default if not set)  
- `docker_image` - new image to assign to container in the deployment, including tag (`drone/drone:latest`)


The following is a sample k8s deployment configuration in your `.drone.yml` file:

```yaml
deploy:
  k8s-deployment:
    image: peloton/drone-k8s-deployment
    url: https://k8s.server/
    token: asldkfj
    insecure: false
    deployment_name: mything
    container_name: mything
    namespace: mynamespace
    docker_image: drone/drone:latest
```

if you want to add secrets for the token it's KUBERNETES_TOKEN
