CONFIG ?= ./configs/config.yml

config: 
	@echo "config file address:${CONFIG}"
run: 
	@nodemon --exec go run ./cmd/server/goo.go --signal SIGTERM

build: 
	@go build -o bin/goo cmd/server/goo.go 

clean: 
	@rm -rf bin

