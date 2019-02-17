.PHONY: gen test lint example

example: lint
	go run -tags=example cmd/xdg-example/main.go

lint: test
	gometalinter ./...

test: gen
	go test -v --race ./...

gen:
	go generate ./...

