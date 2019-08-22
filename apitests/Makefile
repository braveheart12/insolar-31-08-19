TEST_COUNT ?= 1

.PHONY: all
all: codegen-latest smoke ## run api tests

.PHONY: tests
tests: ## run api tests
	go test ./tests -count=$(TEST_COUNT)

.PHONY: smoke
smoke: ## run api smoke tests
	CGO_ENABLED=1 go test -test.v ./tests/smoketests -count=$(TEST_COUNT)

.PHONY: codegen-latest
codegen-latest: pull codegen ## generate api client from latet repo state

.PHONY: pull
pull: ## get actual state in api respositories
#	chmod +x apiclient/pull-repos.sh
	. scripts/pull-repos.sh

.PHONY: codegen
codegen: ## generate api client
#	chmod +x apiclient/generate-api-cli.sh
	. scripts/generate-api-cli.sh

.PHONY: watch
watch: ## start all web specs on localhost
#	chmod +x apiclient/generate-api-cli.sh
	. scripts/watch-api-frontend.sh
