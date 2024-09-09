include .mk/bumpversion.mk

# Makefile Configuration for opus.bdl.datapool.searchengine

# Global Flags and Shell Configuration
MAKEFLAGS += --warn-undefined-variables --no-print-directory
.SHELLFLAGS := -eu -o pipefail -c
SHELL := bash # Use bash for inline if-statements for better control structures

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	MD5SUM := md5sum
endif
ifeq ($(UNAME_S),Darwin)
	MD5SUM := md5 -r
endif

# Check if debug mode is enabled
DEBUG ?= 0

# Depending on the DEBUG variable, set the prefix to suppress command echo
ifeq ($(DEBUG),1)
    QUIET :=
else
    QUIET := @
endif

# Artifact settings
export APP_NAME = opus.bdl.datapool.searchengine
APP_NAME_DASH := $(subst .,-,$(APP_NAME))
export DOCKER_LOCATION ?= ecm-software.artrepo.svanet.ch
export DOCKER_OWNER ?= sol-auto
export DOCKER_REPOSITORY_ROOT := $(DOCKER_LOCATION)/$(DOCKER_OWNER)
export DOCKER_REPOSITORY := $(DOCKER_REPOSITORY_ROOT)/$(APP_NAME_DASH)
export DOCKER_NETWORK ?= opus.bdl.datapool.infra
export DOCKER_NETWORK_EXTERNAL ?= true

# Testing/Linting
 EXPORT_RESULT = false

# Container Tool
DOCKER ?= docker
DOCKER_COMPOSE ?= docker compose

# Enable BuildKit for Docker build
export DOCKER_BUILDKIT := 1
export BUILDKIT_INLINE_CACHE := 1
export COMPOSE_DOCKER_CLI_BUILD := 1
export DOCKER_BUILDPLATFORM ?= linux/arm64
export DOCKER_TARGETPLATFORM ?= $(DOCKER_BUILDPLATFORM),linux/amd64

# GitLab
export GITLAB_PROJECT_ID := 383
export GITLAB_URL_IGS ?= https://gitlab.svanet.ch

# Project config
DEFAULT_PORT := 8000
RANDOM_PORT = $(shell DEFAULT_PORT=$(DEFAULT_PORT); while lsof -i:$$DEFAULT_PORT &>/dev/null; do DEFAULT_PORT=$$(expr $$DEFAULT_PORT + 1); done; echo $$DEFAULT_PORT)

export SWAGGER_VALIDATION ?= false
export AUTH ?= false
export REPO_TYPE ?= redis

# Build-Information
GIT_COMMIT = $(shell git rev-parse --short HEAD)
BUILD_DATE = $(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS = -X 'gitlab.svanet.ch/tibco_bw/opus.bdl.datapool.api/version.GitCommit=$(GIT_COMMIT)' \
          -X 'gitlab.svanet.ch/tibco_bw/opus.bdl.datapool.api/version.BuildDate=$(BUILD_DATE)' \
          -X 'gitlab.svanet.ch/tibco_bw/opus.bdl.datapool.api/version.Version=$(BUMPVERSION)'

# ANTLR
export ANTLR_VERSION := 4.13.2

# Define source grammar and output directories
GRAMMAR_SRC := SearchQuery.g4
GEN_DIR := generated

# Define how to invoke ANTLR, using the version and grammar file
antlr = docker run --rm -it -v $(PWD):/app ${DOCKER_REPOSITORY}-antlr antlr4 -Dlanguage=Go -no-listener -visitor -o /app/generated/$(2) -package $(2) /app/antlr/$(1)

##@ Helpers
# Displays this help message, dynamically generating the command list
help: APP_NAME_LENGTH = $(shell echo -n $(APP_NAME) | wc -c)
help: BOX_WIDTH = 80
help: PADDING_LEFT = $$(($$(($(BOX_WIDTH) - 2 - $(APP_NAME_LENGTH))) / 2))
help: PADDING_RIGHT = $$(($(BOX_WIDTH) - 2 - $(APP_NAME_LENGTH) - $(PADDING_LEFT)))
help: SPACES_LEFT = $(shell printf '%*s' $(PADDING_LEFT) '')
help: SPACES_RIGHT = $(shell printf '%*s' $(PADDING_RIGHT) '')
help: ## Displays this help message
	@echo -e "\033[1;32m╔═══════════════════════════════════════════════════════════════════════════════╗\033[0m"
	@echo -e "\033[1;32m║                                                                               ║\033[0m"
	@echo -e "\033[1;32m║$(SPACES_LEFT)$(APP_NAME)$(SPACES_RIGHT) ║\033[0m"
	@echo -e "\033[1;32m║                                                                               ║\033[0m"
	@echo -e "\033[1;32m╚═══════════════════════════════════════════════════════════════════════════════╝\033[0m"
	@awk 'BEGIN {FS = ":.*##"; printf "\033[36m\033[0m"} /^[a-zA-Z0-9_%\/-]+:.*?##/ { printf "  \033[36m%-50s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
	@printf "\n"

init-delay/%: ## Waits after a delay of N seconds
	@countdown=$(notdir $@); \
	while [ $$countdown -ge 0 ]; do \
		echo -ne "Remaining seconds $$countdown \r"; \
		sleep 1; \
		countdown=$$((countdown - 1)); \
	done; \
	echo -ne "\n"

##@ Docker Commands
docker-login: DOCKER_LOGIN_CREDENTIALS ?= "-u $(DOCKER_USERNAME)"
docker-login: ## Auto login to the Docker repository
	$(QUIET)echo $(DOCKER_PASSWORD) | docker login $(DOCKER_LOGIN_CREDENTIALS) $(DOCKER_REPOSITORY) --password-stdin

docker-builder: ## Create a Docker buildx builder instance if it doesn't already exist
	$(QUIET)$(DOCKER) buildx create --name mybuilder --use 2>/dev/null || true

docker-build: DARGS ?= --load
docker-build: FILE ?= docker-compose.yml
docker-build: PLATFORM ?= $(DOCKER_BUILDPLATFORM)
docker-build: docker-builder ## Build the Docker image with the current version
	@echo $(BUMPVERSION)
	$(DOCKER) buildx bake -f docker-compose.yml $(DARGS) --no-cache --set "*.platform=$(PLATFORM)"

docker-deploy: FILE?=docker-compose.yml
docker-deploy: ## Push the Docker image to the repository
	$(QUIET)$(MAKE) docker-build DARGS=--push FILE=$(FILE) PLATFORM=$(DOCKER_TARGETPLATFORM)

docker-up: ARGS=--build
docker-up: ## Create and start the entire stack
	$(QUIET)$(DOCKER_COMPOSE) up $(ARGS)
du: docker-up ## Alias for docker-up

docker-clean: ## Stops and removes all running Docker containers and their associated volumes
	$(QUIET)docker stop $(shell docker ps -aq) 2>/dev/null || true
	$(QUIET)docker rm -v $(shell docker ps -aq) 2>/dev/null || true
	$(QUIET)$(MAKE) init-delay/3 > /dev/null
	$(QUIET)docker volume rm $(shell docker volume ls -q) 2>/dev/null || true

##@ Go Commands
go-build: ## Build your project and put the output binary in out/bin/
	$(QUIET)mkdir -p out/bin
	$(QUIET)GO111MODULE=on go build -mod vendor -o out/bin/$(APP_NAME) .

go-server: export SERVICE_PORT=$(RANDOM_PORT)
go-server: ## Run the Go application with a random port using go run
	$(QUIET)echo "Running bdl-datapool-api on port $(SERVICE_PORT)..."
	$(QUIET)go run main.go serve

go-clean: ## Remove build related file
	$(QUIET)rm -fr ./bin
	$(QUIET)rm -fr ./out

go-watch: export SERVICE_PORT=$(RANDOM_PORT)
go-watch: ## Run the code with cosmtrek/air to have automatic reload on changes
	$(eval PACKAGE_NAME=$(shell head -n 1 go.mod | cut -d ' ' -f2))
	$(QUIET)$(DOCKER) run -it --rm -w /go/src/$(PACKAGE_NAME) -v $(PWD):/go/src/$(PACKAGE_NAME) -p $(DEFAULT_PORT):$(DEFAULT_PORT) cosmtrek/air

go-test: ## Run the tests of the project
	$(QUIET)go test -v -race ./...

go-coverage: ## Run the tests of the project and export the coverage
	@echo "Running coverage analysis"
	$(QUIET)go test -coverprofile=profile.cov ./...

go-bench: ## Run the benchmarks
	$(QUIET)go test -bench=. -benchmem ./...

lint: lint-go lint-dockerfile lint-yaml ## Run all available linters

lint-dockerfile: ## Lint your Dockerfile
# If dockerfile is present we lint it.
ifeq ($(shell test -e ./docker/Dockerfile && echo -n yes),yes)
	$(eval CONFIG_OPTION = $(shell [ -e $(PWD)/.hadolint.yaml ] && echo "-v $(PWD)/.hadolint.yaml:/root/.config/hadolint.yaml" || echo "" ))
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--format checkstyle" || echo "" ))
	$(eval OUTPUT_FILE = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "| tee /dev/tty > checkstyle-report.xml" || echo "" ))
	$(QUIET)$(DOCKER) run --platform $(DOCKER_BUILDPLATFORM) --rm -i $(CONFIG_OPTION) hadolint/hadolint hadolint $(OUTPUT_OPTIONS) - < ./docker/Dockerfile $(OUTPUT_FILE)
endif

lint-go: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	$(QUIET)$(DOCKER) run --rm --platform $(DOCKER_BUILDPLATFORM) -v $(PWD):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --timeout=65s $(OUTPUT_OPTIONS)

lint-yaml: ## Use yamllint on the yaml file of your projects
	$(QUIET)$(DOCKER) run --rm -it --platform $(DOCKER_BUILDPLATFORM) -v $(PWD):/data cytopia/yamllint -f parsable $(shell git ls-files '*.yml' '*.yaml') $(OUTPUT_OPTIONS)

##@ ANTLR
antlr-generate: ## Generate Go files from the ANTLR grammar.
	$(QUIET)echo "Generating parser and lexer from $(GRAMMAR_SRC)..."
	$(call antlr,SearchQuery.g4,sqparser)
