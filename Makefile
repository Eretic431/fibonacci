proto:
	protoc --go_out=plugins=grpc:internal/fibonacci/proto internal/fibonacci/proto/fibonacci.proto

wire:
	wire ./cmd/fibonacci_service