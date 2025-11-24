.PHONY: up down

# Docker compose file
COMPOSE_FILE := docker-compose.yml

up:
	@if [ -z "$$LICENSE_KEY" ]; then \
		echo "Error: LICENSE_KEY environment variable is required"; \
		echo "Usage: LICENSE_KEY='your_license_key' make up"; \
		exit 1; \
	fi
	@echo "Starting all services..."
	@LICENSE_KEY=$$LICENSE_KEY docker-compose -f $(COMPOSE_FILE) up -d --build --force-recreate
	@echo "Services started."

down:
	@echo "Stopping all services..."
	docker-compose -f $(COMPOSE_FILE) down -v

