# Variables
# APP_NAME := gin-webserver



dev: 
	docker compose  --env-file ./.env up -d --remove-orphans   && swag init -d ./ -o ./docs && air  

dev-compose: 
	docker compose --env-file ./.env up  && swag init -d ./ -o ./docs
dev-build: 
	docker compose --env-file ./.env up --build  && swag init -d ./ -o ./docs

quick: 
	swag init -d ./ -o ./docs && air 

refresh-docs:
	swag init -d ./ -o ./docs
