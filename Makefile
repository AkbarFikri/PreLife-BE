include .env

run:
	go run cmd/app/main.go

migrate_create:
	migrate create -ext sql -dir db/migrations $(name)

migrate_up:
	migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)" -path db/migrations up

migrate_down:
	migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)" -path db/migrations down