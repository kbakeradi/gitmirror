.DEFAULT_GOAL := help
.PHONY: help
help:
	@echo "\n\n------------------------------------------\Make Command Help\n\n"
	@echo "\nGet more detail on SST here: https://docs.serverless-stack.com/packages/cli\n"
	@echo "\nTARGETS:\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

run: ## run directly from go run
	cd ./cmd/git-mirror && go run main.go && cd -
build:	## build command
	cd ./cmd/git-mirror && go build && cd -
run-shell:	## build, run and output to shell command
	make build
	./cmd/git-mirror/git-mirror > run-me.sh
