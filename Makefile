.PHONY: default gen test vendor install

example:
	go run -tags=example cmd/xdg-example/main.go

default:
	echo use gen, test, vendor or install

gen:
	go generate ./...

test:
	go test ./...

vendor:
	dep ensure

install:
	go install ./...
