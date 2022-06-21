.SILENT:
.PHONY:

#=================================
# Variables

mongo_container = SERVICE-GO__mongo
#=================================
# Run service

run:
	@echo Run service
	go run main.go

build-service:
	@echo Build service
	go build && ./service-go
#=================================
# Command for docker

up: up_docker info

info: ps info_domen

up-service: up_docker_all info

rebuild: down build up_docker info

up_docker:
	docker-compose up -d mongodb rabbitmq

up_docker_all:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

# флаг -v удаляет все volume (очищает все данные)
down-clear:
	docker-compose down -v --remove-orphans

build:
	docker-compose build

ps:
	docker-compose ps

build-docker-service:
	docker build . -t service:latest
#=================================
# Into container
mongo_bash:
	docker exec -it $(mongo_container) bash

#=================================
# Info for App
info_domen:
	echo '---------------------------------';
	echo '----------DEV--------------------';
	echo http://localhost:8000
	echo RABBIT-ADMIN http://localhost:15672
	echo '---------------------------------';