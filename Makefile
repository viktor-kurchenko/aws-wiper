SHELL := /bin/bash

fmt:
	gofmt -w ./..

lint:
	docker run -t --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v1.60.3 golangci-lint run -v	

build:
	go build .

all: fmt lint build