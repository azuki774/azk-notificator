CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin

SENDER_CMD=$(CURRENT_DIR)/cmd/sender/
SERVER_CMD=$(CURRENT_DIR)/cmd/server/

.PHONY: build start stop test restart

build:
	cd $(SENDER_CMD) && CGO_ENABLED=0 go build -o $(BIN_DIR)/sender ./...
	cd $(SERVER_CMD) && CGO_ENABLED=0 go build -o $(BIN_DIR)/server ./...

start:
	docker compose -f deployment/compose-local.yml up -d

stop:
	docker compose -f deployment/compose-local.yml down

test:
	docker compose -f deployment/redis.yml down
	docker compose -f deployment/redis.yml up -d
	sleep 3s
	go test -v ./...

restart:
	make stop
	make start
