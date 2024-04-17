.DEFAULT_GOAL := build

GO_ENV = docker-compose --env-file .env run --rm golang

fmt:
	$(GO_ENV) go fmt ./...
.PHONY:fmt

vet:
	$(GO_ENV) go vet ./...
.PHONY:vet

build: vet
	$(GO_ENV) go build
.PHONY:build

shell:
	$(GO_ENV) bash
.PHONY:shell

cmd:
	$(GO_ENV)
.PHONY: cmd

