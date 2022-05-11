HOSTNAME=github.com
NAMESPACE=exasol
NAME=exasol-saas
BINARY=terraform-provider-${NAME}
VERSION=0.2
OS_ARCH=darwin_amd64

default: testacc


# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

openapi: 
	npx @openapitools/openapi-generator-cli generate

build:
	go build -o ${BINARY}

install-local: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

install-deps:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.45.2

lint-fix:
	golangci-lint run --print-issued-lines=false --fix ./...

lint:
	golangci-lint run --print-issued-lines=false ./...