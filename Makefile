BIN_DIR := $(GOPATH)/bin
GOLINT=$(BIN_DIR)/golint
PACKAGES=$(shell go list ./... | grep -v /vendor/)
GOLINT_REPO=github.com/golang/lint/golint
PLATFORMS := windows linux darwin
os = $(word 1, $@)

$(GOLINT):
	go get -u $(GOLINT_REPO)

.PHONY: lint
lint: $(GOLINT)
	golint $(PACKAGES)