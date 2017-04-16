TESTS = $(go list -f '{{range .TestGoFiles}}{{.}} {{end}}' ./...)
ASSETS = bindata.go
COMMIT = $(shell git rev-parse --short HEAD)
BIN = marathon-schema
SCHEMA = schemas/*.json
SCHEMA_PREFIX = schemas/
OSARCH = "linux/386 linux/amd64 darwin/amd64 solaris/amd64"

all: clean deps pack build

install:
	go install ./...

clean:
	-rm -f ./$(BIN) ./$(ASSETS) ./$(BIN)_*

pack: clean
	$(GOPATH)/bin/go-bindata -prefix $(SCHEMA_PREFIX) -o $(ASSETS) $(SCHEMA)

build: pack
	go build -ldflags "-X main.build=$(COMMIT)"

$(TESTS): test

# go list ./... | xargs -n1 go test
test: test-deps
	go test $(TESTS)

deps: release-deps test-deps
	go get github.com/urfave/cli
	go get github.com/xeipuuv/gojsonschema
	go get -u github.com/jteeuwen/go-bindata/...

release: clean deps pack build
	gox -osarch=$(OSARCH) -output="{{.Dir}}_{{.OS}}-{{.Arch}}" ./...

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
