example:
	cd examples && npm i && npm run start

build:
	go build -o hype main.go && chmod +x hype && mv hype /home/harish/go/bin

generate:
	cd scaffolding && npm i && npm run start