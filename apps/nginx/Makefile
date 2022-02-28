.PHONY: help 
.DEFAULT_GOAL := help

# Load all the environment variables from .env
export $(cat .env | xargs)

define BROWSER_PYSCRIPT
import webbrowser, sys

webbrowser.open(sys.argv[1])
endef
export BROWSER_PYSCRIPT

BROWSER := python3 -c "$$BROWSER_PYSCRIPT"

define PRINT_HELP_PYSCRIPT
import re, sys

for line in sys.stdin:
	match = re.match(r'^([/a-zA-Z_-]+):.*?## (.*)$$', line)
	if match:
		target, help = match.groups()
		print("%-30s %s" % (target, help))
endef
export PRINT_HELP_PYSCRIPT

APP_ROOT ?= $(shell 'pwd')

export ENVIRONMENT_OVERRIDE_PATH ?= $(APP_ROOT)/env/Makefile.override.dev
export OVERLAY_PATH ?= k8s/overlays/$(STAGE)

export ENV_FILE ?= .env
export HOST_PORT ?= 80
export STAGE_DEV ?= 'dev'
export STAGE_PROD ?= 'prod'
export CONTINER_PORT ?= 80
export DOCKER_BUILD_FLAGS ?= --no-cache
export DOCKER_BUILD_PATH ?= $(APP_ROOT)
export SOURCE_IMAGE ?= nginx
export DOCKER_FILE ?= $(APP_ROOT)/Dockerfile
export TARGET_IMAGE ?= $(REGISTRY_URL)/$(SOURCE_IMAGE):$(VERSION)

-include $(ENVIRONMENT_OVERRIDE_PATH)
include $(APP_ROOT)/Makefile.app
export ENV_PATH ?= $(APP_ROOT)/.env

build-local: ## build docker image for local use
	@docker build $(DOCKER_BUILD_FLAGS) -t $(SOURCE_IMAGE) --target $(TARGET_STAGE) -f $(DOCKER_FILE) $(DOCKER_BUILD_PATH)

build-dev: ## build docker image for development
	@docker build $(DOCKER_BUILD_FLAGS) -t $(SOURCE_IMAGE) --target $(STAGE_DEV) -f $(DOCKER_FILE) $(DOCKER_BUILD_PATH)

build-prod: ## build docker image for prod env
	@docker build $(DOCKER_BUILD_FLAGS) -t $(SOURCE_IMAGE) --target $(STAGE_PROD) -f $(DOCKER_FILE) $(DOCKER_BUILD_PATH)

docker-build: ## build docker image
	@docker build $(DOCKER_BUILD_FLAGS) -t $(TARGET_IMAGE) --target $(TARGET_STAGE) -f $(DOCKER_FILE) $(DOCKER_BUILD_PATH)

docker-tag:
	@docker tag $(SOURCE_IMAGE) $(TARGET_IMAGE)

docker-push: ## docker push
	@docker push $(TARGET_IMAGE)

docker-login:
	@docker login --username $(DOCKER_USER_NAME) --password-stdin $(DOCKER_HUB_PASSWORD)

run: ## run docker contianer
	@docker run -p $(HOST_PORT):$(CONTINER_PORT) --env-file=$(ENV_FILE) $(SOURCE_IMAGE)


edit-kustomize:
	echo $(ENV_PATH)
	@cp $(OVERLAY_PATH)/kustomization.template.yaml $(OVERLAY_PATH)/kustomization.yaml
	@cp $(ENV_PATH) $(OVERLAY_PATH)/.env
	@cd $(OVERLAY_PATH) && kustomize edit add configmap nginx-config --from-env-file .env

apply:## Kubectl apply using kustomize
	@kustomize build $(OVERLAY_PATH) | kubectl apply -f -
	  
help:
	@python3 -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)
