BASE_STACK = docker compose -f docker-compose.yml

compose-up:
		$(BASE_STACK) up --build -d
.PHONY: compose-up

compose-down:
		$(BASE_STACK) down -v
.PHONY: compose-down

test:
		go test -v -race ./internal/...
.PHONY: test

deps:
		go mod tidy && go mod verify
.PHONY: deps