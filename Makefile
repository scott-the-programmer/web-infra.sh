.PHONY: build
build:
	go mod tidy
	go get

.PHONY: test
test:
	go test ./cloudflare
	go test ./linode
