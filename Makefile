SHELL=/bin/bash
RED=\033[0;31m
GREEN=\033[0;32m
BG_GREY=\033[48;5;237m
YELLOW=\033[38;5;202m
NC=\033[0m # No Color
BOLD_ON=\033[1m
BOLD_OFF=\033[21m
CLEAR=\033[2J

LATEST_VERSION := $(shell cat ./latest-version.txt)

.PHONY: help

help:
	@echo "blablo" automation commands:
	@echo
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Docker

run-go-examples: ## run "examples"
	go run ./go-examples 

go-publish: ## publish blablo
	GOPROXY=proxy.golang.org go list -m github.com/oleksii-honchar/blablo@$(LATEST_VERSION)