GATEWAY_BINARY=gatewayApp
AUTH_BINARY=authApp
MAIL_BINARY=mailApp

up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"
up_build: build_gateway build_auth build_mail
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"
build_gateway:
	@echo "Building gatewat binary..."
	cd gateway-service && env GOOS=linux CGO_ENABLED=0 go build -o ${GATEWAY_BINARY} ./cmd
	@echo "Done!"
build_auth:
	@echo "Building auth binary..."
	cd auth-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd
	@echo "Done!"
build_mail:
	@echo "Building mail binary..."
	cd mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd
	@echo "Done!"
