VERSION=0.0.1
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_TAG=$(shell if [ -z "`git status --porcelain`" ]; then git describe --exact-match --tags HEAD 2>/dev/null; fi)
GIT_TREE_STATE=$(shell if [ -z "`git status --porcelain`" ]; then echo "clean" ; else echo "dirty"; fi)
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif


override LDFLAGS += \
  -X main.version=${VERSION} \
  -X main.buildDate=${BUILD_DATE} \
  -X main.gitCommit=${GIT_COMMIT} \
  -X main.gitTreeState=${GIT_TREE_STATE}

ifneq (${GIT_TAG},)
IMAGE_TAG=${GIT_TAG}
override LDFLAGS += -X main.gitTag=${GIT_TAG}
endif


build-binary:
	CGO_ENABLED=0 GOOS=linux go build \
		-ldflags '${LDFLAGS}' \
		-a -installsuffix cgo -o ./bin/appcenter .

fmt:
	gofmt -l -s -w ./
	goimports -l -w ./

crd: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) paths="./..." output:crd:artifacts:config=config/crd/bases

test-fmt:
	gofmt -l -s ./ | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi
	goimports -l ./ | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

codegen:
	./hack/update-codegen.sh

test-codegen:
	./hack/verify-codegen.sh

test: test-fmt test-codegen
	go test ./...

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.1
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif