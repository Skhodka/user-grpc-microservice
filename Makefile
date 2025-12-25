SHELL := /bin/bash

MAIN_DIR = ./cmd/user/main.go
LOCAL_CONFIG_DIR = ./config/local.yaml

PROTO_FILE = user.proto
PROTO_PATH = ./internal/transport/grpc/proto/${PROTO_FILE}
PROTO_OUT = ./internal/transport/grpc/pb

MIG_PATH = ./migrations

local:
	set -a; \
	. .env; \
	set +a; \
	go run ${MAIN_DIR} --config=${LOCAL_CONFIG_DIR}

build_protoc:
	rm ./internal/transport/grpc/pb/user_grpc.pb.go
	rm ./internal/transport/grpc/pb/user.pb.go
	protoc --go_out=${PROTO_OUT} --go_opt=paths=import --go-grpc_out=${PROTO_OUT} --go-grpc_opt=paths=import $(PROTO_PATH)

migrate_up_all:
	set -a; \
	. .env; \
	set +a; \
	migrate -database "postgres://$$DB_MIG_USER:$$DB_MIG_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=require" -path ${MIG_PATH} up
