.PHONY: pb

pb:
	protoc --go_out=. --go-grpc_out=. ./protos/*proto

up:
	docker-compose up -d

down:
	docker-compose down

build:
	go build -o bin/main cmd/main.go

.PHONY: run

run:
	./bin/main

dev:
	go build -o bin/main cmd/main.go
	docker-compose up -d
	./bin/main