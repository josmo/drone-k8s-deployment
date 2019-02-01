# drone-k8s-deployment

[![Build Status](https://cloud.drone.io/api/badges/josmo/drone-k8s-deployment/status.svg)](https://cloud.drone.io/josmo/drone-k8s-deployment)
[![Go Doc](https://godoc.org/github.com/josmo/drone-k8s-deployment?status.svg)](http://godoc.org/github.com/josmo/drone-k8s-deployment)
[![Go Report](https://goreportcard.com/badge/github.com/josmo/drone-k8s-deployment)](https://goreportcard.com/report/github.com/josmo/drone-k8s-deployment)
[![](https://images.microbadger.com/badges/image/peloton/drone-k8s-deployment.svg)](https://microbadger.com/images/peloton/drone-k8s-deployment "Get your own image badge on microbadger.com")

Drone plugin to update a deployment in k8s. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Versions

This repo is using auto-tag from the drone-docker plugin meaning that
1. master will always publish to 'latest' in docker hub peloton/drone-k8s-deployment
2. tags will follow semver at the 1.0.0+ - initial 0.x.x may have breaking changes

## Binary

Build the binary using `go build`:


### Example

## Usage

Build and deploy from your current working directory:

```
docker run --rm                          \
  -e PLUGIN_URL=<source>                 \
  -e PLUGIN_TOKEN=<token>                \
  -e PLUGIN_CERT=<cert>                  \
  -e PLUGIN_INSECURE=<true>              \
  -e PLUGIN_DEPLOYMENT_NAMES=<deployments> \
  -e PLUGIN_CONTAINER_NAMES=<containers>   \
  -e PLUGIN_NAMESPACES=<namespaces>        \ 
  -e PLUGIN_DOCKER_IMAGE=<image>         \
  -v $(pwd):$(pwd)                       \
  -w $(pwd)                              \
  peloton/drone-k8s-deployment 
```

### Contribution

This repo is setup in a way that if you enable a personal drone server to build your fork it will
 build and publish your image (makes it easier to test PRs and use the image till the contributions get merged)
 
* Build local ```DRONE_REPO_OWNER=josmo DRONE_REPO_NAME=drone-k8s-deployment drone exec```
* on your server just make sure you have DOCKER_USERNAME, DOCKER_PASSWORD, and PLUGIN_REPO set as secrets
