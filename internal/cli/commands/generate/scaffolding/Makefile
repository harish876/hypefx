dev/server:
	air -c air.toml

dev/tailwind:
	npm install && npx tailwindcss -i ./input.css -o ./assets/css/styles.css --watch

dev:
	hype-test build routes && make -j2 dev/tailwind dev/server
