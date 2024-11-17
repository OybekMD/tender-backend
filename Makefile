DB_URL := "postgres://postgres:ebot@localhost:5432/tender?sslmode=disable"

help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Run the service
	go run cmd/main.go
	
# run:
# 	docker-compose up --build

run_db:
	docker-compose up -d db

test:
	go test ./...

swag-gen: ## Generate swagger for api
	swag init -g api/router.go -o api/docs

create-migration: ## Create a new migration
	migrate create -ext sql -dir migrations -seq "$(name)"

migrate-up: ## Apply database migrations
	migrate -path migrations -database "${DB_URL}" -verbose up

migrate-down: ## Rollback database migrations
	migrate -path migrations -database "${DB_URL}" -verbose down

migration-version: ## Check the current migration version
	migrate -database ${DB_URL} -path migrations version 

migrate-dirty: ## Apply all unapplied migrations forcefully
	migrate -path ./migrations/ -database "$(DB_URL)" force 1