include .env
export

.PHONY: help
help: ## display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: tool
tool: ## Install tool.
	@aqua i

.PHONY: gen
gen: ## Generate code.
	@oapi-codegen -generate types -package openapi ./api/openapi.yaml > ./pkg/openapi/types.gen.go
	@oapi-codegen -generate chi-server -package openapi ./api/openapi.yaml > ./pkg/openapi/server.gen.go
	@oapi-codegen -generate client -package openapi ./api/openapi.yaml > ./pkg/openapi/client.gen.go
	@go mod tidy

.PHONY: run
run: ## Run server.
	@go run ./cmd/server/main.go
