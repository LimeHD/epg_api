PROJECTNAME=$(shell basename "$(PWD)")
SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.1.dev

# Go related variables.
GOBASE=$(shell pwd)
# GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
# PID file
PID := /tmp/.$(PROJECTNAME).pid
GO:=$(HOME)/.goenv/bin/goenv exec go
PROJECTDIR=$(GOBASE)
SWAGGERTMP=$(PROJECTDIR)/swaggertmp
SWAGGTMP=$(PROJECTDIR)/swaggtmp

build: get
	@echo "Building $(GOFILES) to ./bin"
	$(GO) build -o bin/$(PROJECTNAME) main.go

get:
	@echo "Download dependencies"
	$(GO) get .

install:
	$(GO) install $(GOFILES)

run:
	$(go) run $(GOFILES)

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
	$(GO) clean

path:
	@echo "Export go path"
	export PATH=$PATH:/usr/local/go/bin

swagger-generate:
	./swag init

swagger-windows:
	swagger.exe serve docs/swagger.json

swagger-linux:
	./swagger serve docs/swagger.json

swagger-prefetch:
	mkdir $(SWAGGERTMP)
	git clone https://github.com/go-swagger/go-swagger $(SWAGGERTMP)

swagger-install: swagger-prefetch
	cd $(SWAGGERTMP)/cmd/swagger && go build -o $(PROJECTDIR)
	rm -rf $(SWAGGERTMP)

go-swag-install:
	mkdir $(SWAGGTMP)
	git clone https://github.com/swaggo/swag $(SWAGGTMP)
	cd $(SWAGGTMP)/cmd/swag && go build -o $(PROJECTDIR)
	rm -rf $(SWAGGTMP)

swagger-init: go-swag-install swagger-install

test:
	go test -json; exit 0

help:
	@echo -e "Usage: make [target] ...\n"
	@echo -e "build        		: Creates a project executable file"
	@echo -e "get          		: Install all dependencies"
	@echo -e "install      		: Compile and install packages and dependencies"
	@echo -e "run          		: Compile and run Go program"
	@echo -e "start        		: Start project from ./bin directory"
	@echo -e "start        		: Remove object files and cached files"
	@echo -e "swagger-init 		: Init swagger componenents"
	@echo -e "swagger-generate 	: Generate swagger docs"
	@echo -e '\nProject name: ' $(PROJECTNAME)
	@echo -e "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo -e "Please report any bug or error to the author."

.PHONY: build get install run watch start stop restart clean help path swagger-linux swagger-windows swagger-install test
