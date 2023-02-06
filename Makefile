.PHONY: all clean

default: all
all: clean test build

build:
	go build -o build/http-server cmd/http-server/main.go

test:
	go test ./... -v

run:
	go run cmd/http-server/main.go

clean:
	rm -rf build

