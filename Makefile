GOBIN := $(shell go env GOPATH)/bin
BIN_NAME := hype-test
example:
	cd examples && npm i && npm run start

cli-tool:
	go build -tags=test -o ${BIN_NAME} && chmod +x ${BIN_NAME} && mv ${BIN_NAME} ${GOBIN}

cli-tool-release:
	go build -o hypefx && chmod +x hypefx && mv hypefx ${GOBIN}

generate:
	cd scaffolding && npm i && npm run start