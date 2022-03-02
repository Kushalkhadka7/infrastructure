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
export OVERLAY_ROOT ?= k8s/overlays
export OVERLAY_PATH ?= k8s/overlays/$(STAGE)

export GIT_COMMIT ?= $(shell git rev-parse HEAD)
export GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)

export ENV_FILE ?= .env
export HOST_PORT ?= 4000
export STAGE_DEV ?= 'dev'
export STAGE_PROD ?= 'prod'
export CONTINER_PORT ?= 4000
export DOCKER_BUILD_FLAGS ?= --no-cache
export DOCKER_BUILD_PATH ?= $(APP_ROOT)
export SOURCE_IMAGE ?= manager
export DOCKER_FILE ?= $(APP_ROOT)/Dockerfile
export TARGET_IMAGE ?= $(REGISTRY_URL)/manager
export TARGET_IMAGE_LATEST ?= $(TARGET_IMAGE):$(SOURCE_IMAGE)-$(GIT_BRANCH)-$(GIT_COMMIT)

-include $(ENVIRONMENT_OVERRIDE_PATH)
include $(APP_ROOT)/Makefile.app
export ENV_PATH ?= $(APP_ROOT)/.env


define kustomize-image-edit
	cd $(OVERLAY_PATH) && kustomize edit set image crkushal7/manager=$(1) && cd $(APP_ROOT)
endef

docker-build: ## build docker image
	@docker build $(DOCKER_BUILD_FLAGS) -t $(SOURCE_IMAGE) --target $(STAGE) -f $(DOCKER_FILE) $(DOCKER_BUILD_PATH)

docker-tag:
	@docker tag $(SOURCE_IMAGE) $(TARGET_IMAGE_LATEST)

docker-push: ## docker push
	@docker push $(TARGET_IMAGE_LATEST)

docker-login:
	@docker login --username $(DOCKER_USER_NAME) --password-stdin $(DOCKER_HUB_PASSWORD)

run: ## run docker contianer
	@docker run -p $(HOST_PORT):$(CONTINER_PORT) --env-file=$(ENV_FILE) $(SOURCE_IMAGE)

edit-image-name: edit-kustomize ## edit image name in kustomize
	@$(call kustomize-image-edit,$(TARGET_IMAGE_LATEST))

edit-kustomize:
	echo $(ENV_PATH)
	@cp $(OVERLAY_ROOT)/kustomization.template.yaml $(OVERLAY_PATH)/kustomization.yaml
	@cp $(ENV_PATH) $(OVERLAY_PATH)/.env
	@cd $(OVERLAY_PATH) && kustomize edit add configmap auth-config --from-env-file .env

apply:## Kubectl apply using kustomize
	@kustomize build $(OVERLAY_PATH) | kubectl apply -f -


build-local: ## build docker image for local use
	@docker build $(DOCKER_BUILD_FLAGS) -t $(SOURCE_IMAGE) --target $(TARGET_STAGE) -f $(DOCKER_FILE) $(DOCKER_BUILD_PATH)

build-dev: ## build docker image for development
	@docker build $(DOCKER_BUILD_FLAGS) -t $(SOURCE_IMAGE) --target $(STAGE_DEV) -f $(DOCKER_FILE) $(DOCKER_BUILD_PATH)

build-prod: ## build docker image for prod env
	@docker build $(DOCKER_BUILD_FLAGS) -t $(SOURCE_IMAGE) --target $(STAGE_PROD) -f $(DOCKER_FILE) $(DOCKER_BUILD_PATH)
	  
help:
	@python3 -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)
