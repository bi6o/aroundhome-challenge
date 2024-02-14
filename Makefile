install:
	go get -u github.com/pressly/goose/cmd/goose
	npm install -g @redocly/openapi-cli
	go get -u github.com/joho/godotenv/cmd/godotenv

migrate:
	godotenv -f .env goose -dir ./migrations  -allow-missing up
