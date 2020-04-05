PROJECTNAME=$(shell basename "$(PWD)")
SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.1.dev

# Go related variables.
GOBASE=$(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
# PID file
PID := /tmp/.$(PROJECTNAME).pid

build:
	@echo "Building $(GOFILES) to ./bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/$(PROJECTNAME) $(GOFILES)

get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

start:
	@echo "Starting bin/$(PROJECTNAME)"
	@./bin/$(PROJECTNAME) & echo $(PROJECTNAME) > $(PID)

stop:
	@echo "Stopping bin/$(PROJECTNAME) if it's running"
	@echo "PID file: \n" $(PID)
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)

restart: clear stop clean build start

clear:
	@clear

clean:
	@echo "Cleaning"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

go-path:
	export PATH=$PATH:/usr/local/go/bin

proto:
	protoc  -I . proto/epg.proto --go_out=plugins=grpc:.

.PHONY: build get install run watch start stop restart clean help proto

help: ##@Miscellaneous Show this help
	@echo -e "Usage: make [target] ...\n"
	@echo -e "build        : Creates a project executable file"
	@echo -e "get          : Install all dependencies"
	@echo -e "install      : Compile and install packages and dependencies"
	@echo -e "run          : Compile and run Go program"
	@echo -e "start        : Start project from ./bin directory"
	@echo -e "start        : Remove object files and cached files"
	@echo -e '\nProject name: ' $(PROJECTNAME)
	@echo -e "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo -e "Please report any bug or error to the author."