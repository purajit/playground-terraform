PKG_NAME=number

default: build

build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=darwin go build -o terraform-provider-${PKG_NAME} -a -ldflags '-extldflags "-static"'

install:
	GO111MODULE=on go install

init:
	terraform init

apply: init
	terraform apply

.PHONY: install build
