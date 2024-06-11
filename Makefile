include .env

run:
	go run cmd/app/main.go

migrate_create:
	migrate create -ext sql -dir db/migrations $(name)

migrate_up:
	migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)" -path db/migrations up

migrate_down:
	migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)" -path db/migrations down

deploy:
	git stash
	echo "Start pulling from github repository...."
	git pull origin main
	echo "Start build golang executable file...."
	go build -o main cmd/app/main.go
	echo "Restart service..."
	systemctl restart prelife.service