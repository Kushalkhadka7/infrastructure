.PHONY: cluster delete-clusters help
.DEFAULT_GOAL := help

# Load all the environment variables from .env
export $(cat .env | xargs)

define BROWSER_PYSCRIPT
	import os, webbrowser, sys

	try:
		from urllib import pathname2url
	except:
		from urllib.request import pathname2url

	webbrowser.open("file://" + pathname2url(os.path.abspath(sys.argv[1])))
endef
export BROWSER_PYSCRIPT

define PRINT_HELP_PYSCRIPT
import re, sys

for line in sys.stdin:
	match = re.match(r'^([a-zA-Z_-]+):.*?## (.*)$$', line)
	if match:
		target, help = match.groups()
		print("%-20s %s" % (target, help))
endef
export PRINT_HELP_PYSCRIPT

BROWSER := python -c "$$BROWSER_PYSCRIPT"

APP_ROOT ?= $(shell 'pwd')
CLUSTER_CONFIG ?= $(APP_ROOT)/cluster/config.yml
CLUSTER_NAME ?= "infra"
METAL_LB ?= $(APP_ROOT)/metallb

cluster: ## create cluster
	@kind create cluster --name $(CLUSTER_NAME) --config $(CLUSTER_CONFIG)

delete-clusters: ## delete cluster
	@kind delete clusters $(CLUSTER_NAME)

cluster-info: ## get all clusters
	@kind get clusters

use-cluster-config: ## user current cluster config file
	@kubectl config use-context kind-$(CLUSTER_NAME)

cluster-lb: ## deploy metallb overlay network to cluster
	@kubectl apply -f $(METAL_LB)/namespace.yml
	@kubectl apply -f $(METAL_LB)/metallb.yml
	@kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)" 

help:
	@python3 -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)