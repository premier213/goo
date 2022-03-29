CONFIG ?= ./pkg/configs/config.yml
APP_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG))

run: 
	@nodemon --exec go run ./main.go --signal SIGTERM

build: 
	@go build -o bin/goo ./main.go 

clean: 
	@rm -rf bin