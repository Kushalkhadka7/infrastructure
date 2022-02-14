.PHONY: help cluster delete-clusters

# Loads all env variables from .env file.
export $(cat .env | xargs)

define BROWSER_PYSCRIPT
import webbrowser, sys

webbrowser.open(sys.argv[1])
endef
export BROWSER_PYSCRIPT

BROWSER := python3 -c "$$BROWSER_PYSCRIPT"

APP_ROOT ?= $(shell 'pwd')
CLUSTER_CONFIG ?= $(APP_ROOT)/cluster/config.yml
CLUSTER_NAME ?= "infra"
METAL_LB ?= $(APP_ROOT)/metallb


help:
	@python3 -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)

cluster:
	@kind create cluster --name $(CLUSTER_NAME) --config $(CLUSTER_CONFIG)

delete-clusters:
	@kind delete clusters $(CLUSTER_NAME)

cluster-info:
	@kind get clusters

use-cluster-config:
	@kubectl config use-context kind-$(CLUSTER_NAME)

cluster-lb:
	@kubectl apply -f $(METAL_LB)/namespace.yml
	@kubectl apply -f $(METAL_LB)/metallb.yml
	@kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)" 
