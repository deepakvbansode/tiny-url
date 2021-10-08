ifndef GITLAB_TOKEN
$(error GITLAB_TOKEN is not set)
endif

include date.mk

#variable declaration
GO      				= go
V 						= 0
Q 						= $(if $(filter 1,$V),,@)
DOCKER_BUILD_ARGS		=--build-arg GITLAB_TOKEN=${GITLAB_TOKEN} --build-arg GITLAB_ID=${GITLAB_ID}
PACKAGE  				= ${shell pwd | rev | cut -f1 -d'/' - | rev}
GITHASH 				= $(shell git rev-parse HEAD)
DOCKER_BUILD_CONTEXT	=.
DOCKER_FILE_PATH		=Dockerfile
BIN      				= $(GOPATH)/bin
DATE    				?= $(shell date +%Y-%m-%d_%I:%M:%S%p)

.DEFAULT_GOAL := all
.PHONY: all
all: fmt concom

#fmt
.PHONY: fmt
fmt: ; $(info $(M) running gofmt) @ ## Run gofmt on all source files
	$Q $(GO) fmt ./...

#to enforce conventional commits
.PHONY: concom
concom:
	@go get github.com/talos-systems/conform
	@chmod +x test/concom.sh
	@test/concom.sh
	@echo "ok"

docker: all ; $(info $(M) building container...) @ ## Build docker container
	@docker build $(DOCKER_BUILD_ARGS) -t $(PACKAGE):$(GITHASH) $(DOCKER_BUILD_CONTEXT) -f $(DOCKER_FILE_PATH)
	@docker tag $(PACKAGE):$(GITHASH) $(PACKAGE):latest

.PHONY: build
build: $(BIN) ; $(info $(M) building executable) @ ## Build program binary
	$Q $(GO) build -tags release -ldflags '-X main.GitComHash=$(GITHASH) -X main.BuildStamp=$(DATE)' -o bin/application cmd/server/main.go

run: docker ;
	@docker run -it -v '$(PWD)':/bin ${PACKAGE}:$(GITHASH)

.PHONY: lint
lint: ; exit 0

protobuf:
	@echo "Skipped - using \"golang.frontdoorhome.com/software/protos\" import instead."