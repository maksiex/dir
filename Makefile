APP_NAME=backv1

# Local running
run:
	go run cmd/main.go

# building with docker
docker-build:
	docker build -t ${APP_NAME} .

# running with docker
docker-run:
	docker run --rm -p 8080:8080 ${APP_NAME}

# build and running
docker-up:
	docker-compose up --build

# deleting docker image
docker-clean:
	docker rmi ${APP_NAME} || true

# logs
logs:
	docker-compose logs -f backend

# stop containers
stop:
	docker-compose down

# rebuild containers
rebuild:
	docker-compose up --build --force-recreate