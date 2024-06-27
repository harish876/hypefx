GOBIN := $(shell go env GOPATH)/bin
BIN_NAME := hype-test
example:
	cd examples && npm i && npm run start

cli-test:
	go build -tags=test -o ${BIN_NAME} && chmod +x ${BIN_NAME} && mv ${BIN_NAME} ${GOBIN}

cli-release:
	go build -o hypefx && chmod +x hypefx && mv hypefx ${GOBIN}

.PHONY: test-cli

test-cli:
	@mkdir foobar
	cd foobar && hype-test generate foobar && go mod tidy
