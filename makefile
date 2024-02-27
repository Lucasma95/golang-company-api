# Usage:
#	make mockery-install
mockery-install:
	@go install github.com/vektra/mockery/v2@v2.20.2

# Usage:
#	make swag
swag:
	$$GOPATH/bin/swag init --parseDependency --parseInternal --parseDepth 2 -o ./docs/swagger -g ./src/api/application/company-api.go

# Usage:
#	make build-mocks
build-mocks:
	mockery --all --dir src/api/core/providers --output src/api/test/mocks/providers --case snake --disable-version-string
	mockery --all --dir src/api/core/usecases --output src/api/test/mocks/usecases --case snake --disable-version-string

# Usage:
#	make test
test:
	@go test ./src/api/...

# Usage:
#	make linter-install
linter-install:
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2

# Usage:
#	make lint
lint:
	@~/go/bin/golangci-lint run

# Usage:
#	make fmt
fmt:
	@go fmt ./...

# Usage:
#	make goimports
goimports:
	@goimports -w src/ap

# Usage:
#	make local-infra
make local-infra:
	@echo "Creating local infrastructure"
	@docker compose -f docker-compose.yml up -d
