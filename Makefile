example:
	cd examples && npm i && npm run start

build:
	go build -o hype main.go && mv hype /Users/harishgokul/go/bin

generate:
	cd scaffolding && npm i && npm run start