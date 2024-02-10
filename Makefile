# Define Docker compose command
DOCKER_COMPOSE = docker-compose

# Define the name of your Docker compose file
COMPOSE_FILE = docker-compose.yml

# Define the name of your app service
APP_SERVICE = app

# Define targets
.PHONY: help build up down logs clean

# Display help message
help:
	@echo "Usage: make [TARGET]"
	@echo ""
	@echo "Targets:"
	@echo "  build      Build Docker images"
	@echo "  up         Start Docker containers"
	@echo "  down       Stop and remove Docker containers"
	@echo "  logs       View container logs"
	@echo "  clean      Remove Docker volumes and images"

# Build Docker images
build:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) build

# Start Docker containers
up:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up -d

# Stop and remove Docker containers
down:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down

# View container logs
logs:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) logs -f $(APP_SERVICE)

# Remove Docker volumes and images
clean:
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down -v --remove-orphans
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) rm -f -v
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down --rmi all --volumes --remove-orphans
