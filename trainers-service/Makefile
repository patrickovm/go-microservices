.PHONY: pb

pb:
	protoc --go_out=. --go-grpc_out=. ./protos/*proto

build:
	go build -o bin/main cmd/main.go

run:
	./bin/main

up:
	docker-compose up -d

down:
	docker-compose down

dev:
	go build -o bin/main cmd/main.go
	docker-compose up -d
	./bin/main