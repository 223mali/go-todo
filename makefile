# Variables
# APP_NAME := gin-webserver


dev: 
	docker compose  --env-file ./.env up -d  && swag init -d ./src -o ./src/docs && air  

quick: 
	swag init -d ./src -o ./src/docs && air 

refresh-docs:
	swag init -d ./src -o ./src/docs
