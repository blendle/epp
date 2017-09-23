BINARY=bin/epp

VERSION=$(shell git for-each-ref refs/tags --sort=-committerdate --format='%(refname:short)' --count=1)
GIT_COMMIT=`git rev-parse @`
LDFLAGS=-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT)

build:
	mkdir -p bin
	go build -ldflags "$(LDFLAGS)" -o $(BINARY)

release:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w $(LDFLAGS)" -o $(BINARY)
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w $(LDFLAGS)" -o $(BINARY)_darwin64
