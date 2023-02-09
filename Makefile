.PHONY: all build test coverage run clean

default: all
all: clean test build

clean:
	go clean
	rm -rf build

build:
	go build -o build/gca-server cmd/http-server/main.go
	go build -o build/gca-cli cmd/cli/main.go

test:
	go test ./... -v

coverage:
	go test ./... -coverprofile=build/coverage.out > /dev/null
	go tool cover -func build/coverage.out

run:
	go run cmd/http-server/main.go



