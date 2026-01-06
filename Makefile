BASE_STACK = docker compose -f docker-compose.yml

compose-up: ### Run docker-compose
		$(BASE_STACK) up --build -d
.PHONY: compose-up

compose-down: ### Down docker-compose
		$(BASE_STACK) down -v
.PHONY: compose-down

run: deps ### local run
		go run ./cmd/app/main.go
.PHONY: run

test: ### run tests
		go test -v -race ./internal/...
.PHONY: test

deps: ### deps tidy + verify
		go mod tidy && go mod verify
.PHONY: deps