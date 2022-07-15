# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

compose-up: ### Run docker-compose
	docker-compose -f docker-compose-dev.yml up --build -d postgres
.PHONY: compose-up

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

ent: ### Generate ent assets and code
	go generate ./ent
.PHONY: ent

docs: ### Generate swagger docs
	swag init -g cmd/ektimo-api/main.go
.PHONY: docs

run: docs ### Run ektimo-api locally
	go mod tidy && go mod download && \
	ENV=dev CGO_ENABLED=0 go run ./cmd/ektimo-api
.PHONY: run

test: ### Run tests
	go test -v -cover -race ./internal/...
.PHONY: test