CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin

SENDER_CMD=$(CURRENT_DIR)/cmd/sender/
SERVER_CMD=$(CURRENT_DIR)/cmd/server/

.PHONY: build

build:
	cd $(SENDER_CMD) && CGO_ENABLED=0 go build -o $(BIN_DIR)/sender ./...
	cd $(SERVER_CMD) && CGO_ENABLED=0 go build -o $(BIN_DIR)/server ./...
