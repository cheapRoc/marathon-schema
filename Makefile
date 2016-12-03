TESTS = $(go list -f '{{range .TestGoFiles}}{{.}} {{end}}' ./...)
ASSETS = bindata.go
COMMIT = $(shell git rev-parse --short HEAD)
BIN = marathon-schema
SCHEMAS = schemas/AppDefinition.json schemas/Group.json

all: clean deps pack build test

install:
	go install ./...

clean:
	-rm -f $(BIN) $(ASSETS) $(BIN)_*

pack: clean
	go-bindata -o $(ASSETS) $(SCHEMAS)

build: pack
	go build -ldflags "-X main.Build=$(COMMIT)"

$(TESTS): test

test: test-deps
	go test $(TESTS)
	# go list ./... | xargs -n1 go test

deps:


release: clean pack deps build test
	gox ./...
	# upload to yonder cloud

test-deps:
	go get github.com/stretchr/testify

release-deps:
	go get github.com/mitchellh/gox

list-deps:
	go list -f '{{range .Deps}}{{.}} {{end}}' ./... | tr ' ' \\n

list-imports:
	go list -f '{{range .Imports}}{{.}} {{end}}' ./... | tr ' ' \\n

list-test-imports:
	go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | tr ' ' \\n


.PHONY: clean build
