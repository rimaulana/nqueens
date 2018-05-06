BINARY := nqueens
VERSION ?= vlatest
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
	for PKG in $(PACKAGES); do \
		golint -set_exit_status $$PKG || exit 1; \
	done;

.PHONY: test
test: clean
	echo "mode: count" > coverage.out; \
	for PKG in $(PACKAGES); do \
		go test -v -covermode=count -coverprofile=profile.out $$PKG; \
		if [ -f profile.out ]; then \
        	cat profile.out | grep -v "mode:" >> coverage.out; \
        	rm profile.out; \
    	fi; \
	done;

.PHONY: cover
cover: test
	go tool cover -html=coverage.out -o=coverage.html; \
	rm coverage.out;

.PHONY: clean
clean:
	for COV in $(shell ls | grep coverage); do \
		rm $$COV; \
	done; \
	rm -rf release;

.PHONY: $(PLATFORMS)
$(PLATFORMS): clean
	mkdir -p release; \
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64;

.PHONY: release
release: windows linux darwin