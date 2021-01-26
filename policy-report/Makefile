# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: manifests build fmt vet

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) paths="./..." output:dir=crd
	$(CONTROLLER_GEN) object paths=./api/v1alpha1/

# Run go build against code
build:
	go build ./...

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.6 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

# generate will create api docs
generate:
	gen-crd-api-reference-docs -api-dir ./api/v1alpha1 -config docs/config.json -template-dir docs/template -out-file docs/index.html
	# TODO - Currently gen-crd-api-reference-docs doesn't support kubebuilder for generating docs. https://github.com/ahmetb/gen-crd-api-reference-docs/issues/15
	# go run github.com/ahmetb/gen-crd-api-reference-docs -api-dir ./api/v1alpha1 -config docs/config.json -template-dir docs/template -out-file docs/index.html

