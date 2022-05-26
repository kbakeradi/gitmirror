.DEFAULT_GOAL := help
.PHONY: help
help:
	@echo "\n\n------------------------------------------\Make Command Help\n\n"
	@echo "\nGet more detail on SST here: https://docs.serverless-stack.com/packages/cli\n"
	@echo "\nTARGETS:\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

run:
	cd cmd/git-mirror && go run main.go && cd -
