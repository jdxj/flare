ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"
IMAGE_TAG = jdxj/flare:v0.1.0

include ./hack/hack-cli.mk
include ./hack/hack.mk

.PHONY: image.build
image.build:
	@docker build -t $(IMAGE_TAG) -f manifest/docker/Dockerfile .

.PHONY: image.push
image.push: image.build
	@docker push $(IMAGE_TAG)
