APP_NAME := go-microshop
APP_PATH := ./cmd/api
CONFIG_DIR := ./config
IMAGE := ghcr.io/jkhelil/$(APP_NAME)
TAG := latest

KO := ko
GO := go
GOLANGCI_LINT := golangci-lint

.PHONY: all build lint-go test e2e deploy undeploy clean

all: build

vendor:
	@echo "📦 Vendoring dependencies..."
	$(GO) mod tidy
	$(GO) mod vendor
	@echo "✅ Dependencies vendored into ./vendor"

build: vendor
	@echo "🚧 Building $(APP_NAME)..."
	$(GO) build -buildvcs=false -mod=vendor -o bin/$(APP_NAME) $(APP_PATH)
	@echo "✅ Build complete: bin/$(APP_NAME)"

lint-go:
	@echo "Linting go files..."
	@$(GOLANGCI_LINT) run ./pkg/... ./e2e/... --modules-download-mode=vendor \
							--max-issues-per-linter=0 \
							--max-same-issues=0 \
							--timeout 10m

test:
	@echo "🧪 Running unit tests..."
	$(GO) test -v ./... -short -cover

e2e:
	@echo "🌐 Running e2e tests..."
	$(GO) test -v ./tests/e2e/...

deploy:
	@echo "🚀 Deploying $(APP_NAME) using ko..."
	KO_DOCKER_REPO=$(IMAGE) $(KO) apply -f $(CONFIG_DIR)
	@echo "✅ Deployment applied using ko"

undeploy:
	@echo "🧹 Removing deployed resources..."
	kubectl delete -f $(CONFIG_DIR) --ignore-not-found=true
	@echo "✅ Resources cleaned up"

clean:
	@echo "🧼 Cleaning local build..."
	rm -rf bin
