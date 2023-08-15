docker-up:
	docker compose up

docker-down:
	docker compose down 

docker-build-api-gateway: ## build api gateway service docker image
	docker compose build api-gateway

dokcker-build-auth-service:## build auth service docker image
	docker compose build auth-service

docker-build-user-service: ## build api user service docker image
	docker compose build user-service

docker-build-user-service: ## build api product service docker image
	docker compose build product-service

docker-build-cart-service: ## build api cart service docker image
	docker compose build cart-service


