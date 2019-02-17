.PHONY: default gen test example

default:
	echo use gen, test, vendor or install

gen:
	go generate ./...

test:
	go test ./...

example:
	go run -tags=example cmd/xdg-example/main.go

