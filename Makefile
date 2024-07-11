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

PLATFORMS ?= linux/amd64 darwin/arm64
.PHONY: haul
haul: hauler ## Generate Hauls for latest release
	$(foreach platform,$(PLATFORMS),\
		$(eval GOOS=$(word 1,$(subst /, ,$(platform)))) \
		$(eval GOARCH=$(word 2,$(subst /, ,$(platform)))) \
		echo "Building Haul for $(GOOS)/$(GOARCH)..."; \
		hauler store sync -s store-$(GOOS)-$(GOARCH) -f hauler-manifest.yaml -p $(platform); \
		hauler store save -s store-$(GOOS)-$(GOARCH) -f validator-haul-$(GOOS)-$(GOARCH).tar.zst; \
		rm -rf store-$(GOOS)-$(GOARCH);

HAULER_VERSION ?= 1.0.4
.PHONY: hauler
hauler: ## Install hauler
	curl -sfL https://get.hauler.dev | HAULER_VERSION=$(HAULER_VERSION) bash