# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: manifests build fmt vet

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) crd paths="./pkg/api/..." output:dir=crd/v1beta2
	$(CONTROLLER_GEN) object paths="./pkg/api/..."

# Run go build against code
build:
	go build ./...

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

CONTROLLER_TOOLS_VERSION ?= v0.14.0
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen-$(CONTROLLER_TOOLS_VERSION)

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	$(call go-install-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen,$(CONTROLLER_TOOLS_VERSION))

# generate will create api docs
generate:
	gen-crd-api-reference-docs -api-dir ./pkg/api/reports.x-k8s.io/v1beta2 -config ./docs/config.json -template-dir ./docs/template -out-file ./docs/index.html
	# gen-crd-api-reference-docs -api-dir ./pkg/api/v1beta1 -config docs/config.json -template-dir docs/template -out-file docs/index.html
	# TODO - Currently gen-crd-api-reference-docs doesn't support kubebuilder for generating docs. https://github.com/ahmetb/gen-crd-api-reference-docs/issues/15
	# go run github.com/ahmetb/gen-crd-api-reference-docs -api-dir ./api/v1alpha1 -config docs/config.json -template-dir docs/template -out-file docs/index.html

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# $1 - target path with name of binary (ideally with version)
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f $(1) ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv "$$(echo "$(1)" | sed "s/-$(3)$$//")" $(1) ;\
}
endef
