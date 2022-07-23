include .env
export

POSTGRES_CONNECTION_STRING := postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):5432/$(PG_DATABASE)?sslmode=disable

clean:
	rm -rf ./bin

build:
	go build -o bin/main main.go

run:
	go run main.go

compile:
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

migrate_up:
	migrate -source file://database/migrations -database $(POSTGRES_CONNECTION_STRING) up

migrate_down:
	migrate -source file://database/migrations -database $(POSTGRES_CONNECTION_STRING) down

migrate_new:
	migrate create -ext sql -dir database/migrations $(name)
