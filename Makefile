include build/makelib/common.mk
include build/makelib/plugin.mk

# Image URL to use all building/pushing image targets
IMG ?= quay.io/validator-labs/validator:latest
CERTS_INIT_IMG ?= quay.io/validator-labs/validator-certs-init:latest

# Helm vars
CHART_NAME=validator

VALIDATION_RESULTS_CRD = chart/validator/crds/validation.spectrocloud.labs_validationresults.yaml
VALIDATOR_CONFIGS_CRD = chart/validator/crds/validation.spectrocloud.labs_validatorconfigs.yaml

reviewable-ext:
	@$(INFO) Checking for plugin version updates...
	bash hack/update-versions.sh
	rm $(VALIDATION_RESULTS_CRD) $(VALIDATOR_CONFIGS_CRD)
	cp config/crd/bases/validation.spectrocloud.labs_validationresults.yaml $(VALIDATION_RESULTS_CRD)
	cp config/crd/bases/validation.spectrocloud.labs_validatorconfigs.yaml $(VALIDATOR_CONFIGS_CRD)

.PHONY: docker-build-certs-init
docker-build-certs-init: ## Build validator-certs-init docker image.
	$(CONTAINER_TOOL) build -f hack/validator-certs-init.Dockerfile -t ${CERTS_INIT_IMG} . --platform linux/$(GOARCH)

HAUL_PLATFORMS ?= linux/amd64 linux/arm64
.PHONY: haul
haul: hauler ## Generate Hauls for latest release
	$(foreach platform,$(HAUL_PLATFORMS),\
		$(eval GOOS=$(word 1,$(subst /, ,$(platform)))) \
		$(eval GOARCH=$(word 2,$(subst /, ,$(platform)))) \
		echo "Building Haul for $(GOOS)/$(GOARCH)..."; \
		$(HAULER) store sync -s store-$(GOOS)-$(GOARCH) -f hauler-manifest.yaml -p $(platform); \
		$(HAULER) store save -s store-$(GOOS)-$(GOARCH) -f validator-haul-$(GOOS)-$(GOARCH).tar.zst; \
		rm -rf store-$(GOOS)-$(GOARCH);)

HAULER_VERSION ?= 1.0.4
.PHONY: hauler
hauler: ## Install hauler
	@command -v hauler >/dev/null 2>&1 || { \
		echo "Hauler version $(HAULER_VERSION) not found, downloading..."; \
		curl -sfL https://get.hauler.dev | HAULER_VERSION=$(HAULER_VERSION) bash; \
	}
HAULER = /usr/local/bin/hauler
