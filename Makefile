# -include will silently skip missing files, which allows us
# to load those files with a target in the Makefile. If only
# "include" was used, the make command would fail and refuse
# to run a target until the include commands succeeded.
-include build/makelib/common.mk

# Image URL to use all building/pushing image targets
IMG ?= quay.io/validator-labs/validator:latest
CERTS_INIT_IMG ?= quay.io/validator-labs/validator-certs-init:latest

.PHONY: docker-build-certs-init
docker-build-certs-init: ## Build validator-certs-init docker image.
	$(CONTAINER_TOOL) build -f hack/validator-certs-init.Dockerfile -t ${CERTS_INIT_IMG} . --platform linux/$(GOARCH)