ROOT_DIRECTORY := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))


.PHONY: build-cli
build-cli: ## Build a CLI binary for the host architecture without embedded UI
	go build -o main.o .

