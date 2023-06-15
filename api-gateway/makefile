proto: 
	protoc --go_out=. --go-grpc_out=. ./pkg/proto/*.proto

mock:
	mockgen -source=./pkg/pb/auth_grpc.pb.go -destination=./pkg/mock/client_mock.go -package=mock

wire:
	cd pkg/di && wire

run: 
	go run cmd/main.go

docker-build:
	docker build -t nikhil382/ecommerce-microservice-clean-arch-api-gateway .