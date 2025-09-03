## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build: builds the command line tool dist directory
build:
	@go build -o ./dist/celeritas ./cmd/cli

build_cli:
	@go build -o ../myapp/celeritas ./cmd/cli