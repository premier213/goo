CONFIG ?= ./config/config.yml
APP_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG))

run: 
	@nodemon --exec go run ./cmd/server/goo.go --signal SIGTERM

build: 
	@go build -o bin/goo cmd/server/goo.go 

clean: 
	@rm -rf bin

migrate-init: 
	@migrate create -ext sql -dir internal/db/migration -seq init_schema
migrate-up:
	migrate -path internal/db/migration -database "${APP_DSN}" -verbose up
migrate-down:
	migrate -path internal/db/migration -database "${APP_DSN}" -verbose down