CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin/

CONTAINER_SENDER:=azk-notificator-sender:latest
CONTAINER_SERVER:=azk-notificator-server:latest
SENDER_CMD=$(CURRENT_DIR)/cmd/sender/
SERVER_CMD=$(CURRENT_DIR)/cmd/server/

.PHONY: build start stop test restart

build:
	CGO_ENABLED=0 go build -o build/bin/ ./...
	docker build -t ${CONTAINER_SENDER} -f build/Dockerfile-sender .
	docker build -t ${CONTAINER_SERVER} -f build/Dockerfile-server .

start:
	docker compose -f deployment/compose-local.yml up -d

stop:
	docker compose -f deployment/compose-local.yml down

test:
	docker compose -f deployment/redis.yml down
	docker compose -f deployment/redis.yml up -d
	sleep 3s
	go test -v ./...
	docker compose -f deployment/redis.yml down

restart:
	make stop
	make start
