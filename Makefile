GOBIN := $(shell go env GOPATH)/bin
BIN_NAME := hype
example:
	cd examples && npm i && npm run start

build:
	go build -o hype main.go && chmod +x hype && mv hype ${GOBIN}

generate:
	cd scaffolding && npm i && npm run start