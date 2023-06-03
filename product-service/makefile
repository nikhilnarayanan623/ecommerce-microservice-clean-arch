proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/proto/*.proto

mock-repo:
	mockgen -source=./pkg/repository/interfaces/product.go -destination=./pkg/mock/product_mock.go -package=mock

wire:
	cd pkg/di && wire

run:
	go run cmd/main.go
