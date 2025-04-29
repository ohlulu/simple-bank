.PHONY: run
run:
	nodemon --exec go run main.go --signal SIGINT

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: install
install:
	@echo "Checking and installing required tools..."
	@if ! command -v docker > /dev/null; then \
		echo "Error: Docker is not installed. Please install Docker first: https://docs.docker.com/get-docker/"; \
		exit 1; \
	else \
		printf "$(GREEN)✓$(NC) Docker is installed.\n"; \
	fi
	@if ! command -v migrate > /dev/null; then \
		echo "Installing golang-migrate..."; \
		brew install golang-migrate; \
	else \
		printf "$(GREEN)✓$(NC) golang-migrate is already installed.\n"; \
	fi
	@if ! command -v sqlc > /dev/null; then \
		echo "Installing sqlc..."; \
		brew install sqlc; \
	else \
		printf "$(GREEN)✓$(NC) sqlc is already installed.\n"; \
	fi
	@printf "$(GREEN)✓$(NC) All required tools are installed or have been installed.\n";

.PHONY: postgres
postgres:
	@# create by
	@# docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.4-alpine
	docker start postgres17

.PHONY: createdb
createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

.PHONY: dropdb
dropdb:
	docker exec -it postgres17 dropdb --username=root simple_bank

.PHONY: migrateup
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
	
.PHONY: migrateup1
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

.PHONY: migratedown
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: migratedown1
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: mock
mock:
	mockgen -destination db/mock/store.go github.com/ohlulu/simple-bank/db/sqlc Store

### ----------------------- Helper ----------------------- ###

NC    	= \033[0m
GREEN 	= \033[0;32m
BLUE 	= \033[0;34m
RED 	= \033[0;31m
YELLOW	= \033[1;33m