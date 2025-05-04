APP_NAME=backv1

# Local running
run:
	go run cmd/main.go

# Build Docker image
docker_build:
	docker build -t ${APP_NAME} .

# Run container
docker_run:
	docker run --rm -p 8080:8080 ${APP_NAME}

# Build and run with docker-compose
docker_up:
	docker compose up --build -d

# Remove Docker image
docker_clean:
	docker rmi ${APP_NAME} || true

# Logs
logs:
	docker compose logs -f backend

# Stop containers
stop:
	docker compose down

# Rebuild containers
rebuild:
	docker compose up --build -d --force-recreate
