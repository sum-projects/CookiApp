AUTH_BINARY=authApp

up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"
up_build: build_auth
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"
build_auth:
	@echo "Building broker binary..."
	cd auth && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd
	@echo "Done!"