GOPACKAGES=$(shell find . -name '*.go' -not -path "$(SOURCE_DIR)/vendor/*" -exec dirname {} \; | uniq)

build:
	go build -o ./bin/service-auth ./cmd/main

test:
	GOPATH=$(shell pwd) go test -v $(GOPACKAGES)
