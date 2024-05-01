GOBIN := $(shell go env GOPATH)/bin
BIN_NAME := hype-test
example:
	cd examples && npm i && npm run start

test-build:
	go build -o ${BIN_NAME} main.go && chmod +x ${BIN_NAME} && mv ${BIN_NAME} ${GOBIN}

generate:
	cd scaffolding && npm i && npm run start