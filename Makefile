.PHONY: install
install:
	brew install golang-migrate
	brew install sqlc

.PHONY: postgres
postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.4-alpine

.PHONY: createdb
createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

.PHONY: dropdb
dropdb:
	docker exec -it postgres17 dropdb --username=root simple_bank

.PHONY: migrateup
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

.PHONY: migratedown
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down



