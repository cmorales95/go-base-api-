.SILENT:
.EXPORT_ALL_VARIABLES:

# Commands

# Build all the containers and run it from docker
docker-build:
	docker-compose up --build -d

# Destroy all the containers
docker-down-all:
	docker-compose down -v --remove-orphans

# Run only the external services (recommended for development env)
docker-run-services:
	docker-compose up postgres

# Go to the bash
docker-postgres:
	docker exec -it postgres_aug7 bash

docker-remove-orphans-images:
	docker rmi $(docker images -f dangling=true -q)
