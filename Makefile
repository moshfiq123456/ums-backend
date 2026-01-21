.PHONY: help build run test clean fmt lint install-deps docker-up docker-down \
        migrate-up migrate-down migrate-force migrate-create install-migrate debug-db

# -------------------------------
# App Variables
# -------------------------------
APP_NAME=go-training-api
MAIN_PATH=cmd/api/main.go
BINARY_NAME=bin/$(APP_NAME)
GO=go
DOCKER_COMPOSE=docker-compose

# -------------------------------
# Load .env safely (no export)
# -------------------------------
ENV_FILE=.env

define get_env
$(shell sed -n 's/^$(1)=//p' $(ENV_FILE))
endef

DB_USER     := $(call get_env,DB_USER)
DB_PASSWORD := $(call get_env,DB_PASSWORD)
DB_HOST     := $(call get_env,DB_HOST)
DB_PORT     := $(call get_env,DB_PORT)
DB_NAME     := $(call get_env,DB_NAME)
DB_SSLMODE  := $(call get_env,DB_SSLMODE)

# -------------------------------
# Database
# -------------------------------
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATIONS_PATH=migrations

# -------------------------------
# Migration Tool
# -------------------------------
MIGRATE_BIN=$(shell command -v migrate 2> /dev/null)

install-migrate:
ifndef MIGRATE_BIN
	@echo "Installing golang-migrate..."
	curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
	sudo mv migrate /usr/local/bin/
else
	@echo "golang-migrate already installed"
endif

# -------------------------------
# Debug
# -------------------------------
debug-db:
	@echo "DB_USER=$(DB_USER)"
	@echo "DB_HOST=$(DB_HOST)"
	@echo "DB_PORT=$(DB_PORT)"
	@echo "DB_NAME=$(DB_NAME)"
	@echo "DATABASE_URL=$(DB_URL)"

# -------------------------------
# Migrations
# -------------------------------
migrate-up: install-migrate
	@echo "Running migrations up..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down: install-migrate
	@echo "Rolling back last migration..."
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-force: install-migrate
	@read -p "Enter version to force: " version; \
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $$version

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name

# -------------------------------
# Help
# -------------------------------
help:
	@echo ""
	@echo "Available commands:"
	@echo "  make build           - Build the application"
	@echo "  make run             - Run the application"
	@echo "  make dev             - Run with hot reload"
	@echo "  make test            - Run tests"
	@echo "  make fmt             - Format code"
	@echo "  make lint            - Run linter"
	@echo "  make clean           - Clean build artifacts"
	@echo "  make docker-up       - Start PostgreSQL container"
	@echo "  make docker-down     - Stop PostgreSQL container"
	@echo "  make migrate-up      - Run DB migrations"
	@echo "  make migrate-down    - Rollback last migration"
	@echo "  make migrate-create  - Create new migration"
	@echo "  make debug-db        - Print DB config"
	@echo ""

# -------------------------------
# App Commands
# -------------------------------
install-deps:
	$(GO) mod download
	$(GO) mod tidy

build: install-deps
	$(GO) build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BINARY_NAME)"

run: build
	./$(BINARY_NAME)

dev:
	@command -v air >/dev/null 2>&1 || \
	(echo "Installing air..." && $(GO) install github.com/cosmtrek/air@latest)
	air

test:
	$(GO) test -v -cover ./...

fmt:
	$(GO) fmt ./...

lint:
	@command -v golangci-lint >/dev/null 2>&1 || \
	(echo "Installing golangci-lint..." && \
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

clean:
	rm -f $(BINARY_NAME)
	$(GO) clean

# -------------------------------
# Docker
# -------------------------------
docker-up:
	$(DOCKER_COMPOSE) up -d postgres

docker-down:
	$(DOCKER_COMPOSE) down
