proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/proto/*.proto

mock-repo:
	mockgen -source=./pkg/repository/interfaces/product.go -destination=./pkg/mock/product_mock.go -package=mock

wire:
	cd pkg/di && wire

run:
	go run cmd/api/main.go

docker-build:
	docker build -t nikhil382/ecommerce-microservice-clean-arch-product-service .
