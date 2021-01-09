HASH := $(shell git rev-parse --short HEAD)
COMMIT_DATE := $(shell git show -s --format=%ci ${HASH})
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
VERSION := ${HASH} (${COMMIT_DATE})

.PHONY: run
run:
	@go run -race ./cmd/mouser

.PHONY: test
test:
	@go test -race ./...

.PHONY: clean
clean:
	@rm -rf ./bin/*

.PHONY: build
build: clean
	@mkdir -p ./bin
	@go build \
		-o ./bin/ \
		-ldflags="-X 'main.version=${VERSION}' -X 'main.dateate=${BUILD_DATE}'" \
		./cmd/mouser