proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/proto/*.proto

wire:
	cd pkg/di && wire

run:
	go run cmd/api/main.go
	
docker-build:
	docker build -t nikhil382/ecommerce-microservice-clean-arch-cart-service .