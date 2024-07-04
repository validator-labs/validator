include build/makelib/common.mk
include build/makelib/plugin.mk

# Image URL to use all building/pushing image targets
IMG ?= quay.io/validator-labs/validator:latest
CERTS_INIT_IMG ?= quay.io/validator-labs/validator-certs-init:latest

# Helm vars
CHART_NAME=validator

.PHONY: docker-build-certs-init
docker-build-certs-init: ## Build validator-certs-init docker image.
	$(CONTAINER_TOOL) build -f hack/validator-certs-init.Dockerfile -t ${CERTS_INIT_IMG} . --platform linux/$(GOARCH)