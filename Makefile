# ____________________________________________________________________________
# Main
run:
	go run ./cmd/fibonacci_service/

# ____________________________________________________________________________
# Code generation
proto:
	protoc --go_out=plugins=grpc:internal/fibonacci/proto internal/fibonacci/proto/fibonacci.proto

wire:
	wire ./cmd/fibonacci_service

# ____________________________________________________________________________
# Docker
docker-build:
	docker build -f ./docker/Dockerfile . --tag ghcr.io/eretic431/fibonacci

docker-run:
	docker run -p 8080:8080 ghcr.io/eretic431/fibonacci

docker-push: docker-build
	docker push ghcr.io/eretic431/fibonacci:latest

# ____________________________________________________________________________
# Deployments
up:
	docker-compose up -d

down:
	docker-compose down