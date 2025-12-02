# Makefile for miniGame (Docker-based)

COMPOSE := docker compose -f docker-compose.yml

.PHONY: help build up down restart logs ps shell-mysql shell-server test-api clean

help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  build        Build all images (server, web)"
	@echo "  up           Build and start services in background (docker compose up -d --build)"
	@echo "  down         Stop and remove containers (preserves volumes)"
	@echo "  restart      Restart the compose stack"
	@echo "  logs         Tail server logs (pass SERVICE=server|web|mysql)"
	@echo "  ps           List running containers for the project"
	@echo "  shell-mysql  Open a mysql client in the mysql container"
	@echo "  shell-server Exec a shell inside server container"
	@echo "  test-api     Submit a sample score and show recent ranks (requires curl or PowerShell)"
	@echo "  clean        Remove containers and volumes (DANGEROUS: deletes DB data)"

build:
	@echo "Building images..."
	$(COMPOSE) build --no-cache

up:
	@echo "Bringing up services..."
	$(COMPOSE) up -d --build

down:
	@echo "Stopping services..."
	$(COMPOSE) down

restart: down up

logs:
	@if [ "$(SERVICE)" = "" ]; then \
		echo "Please set SERVICE=server or SERVICE=web or SERVICE=mysql"; exit 1; \
	fi
	$(COMPOSE) logs -f $(SERVICE)

ps:
	$(COMPOSE) ps

shell-mysql:
	$(COMPOSE) exec mysql sh -c "mysql -uroot -p$$MYSQL_ROOT_PASSWORD"

shell-server:
	$(COMPOSE) exec server sh

test-api:
	@echo "Posting a sample record to http://localhost:8888/api/records/submit"
	@if command -v curl >/dev/null 2>&1; then \
		curl -s -X POST http://localhost:8888/api/records/submit -H 'Content-Type: application/json' -d '{"name":"make-test","score":123,"comment":"from-make"}' | jq || true; \
	else \
		echo "curl not found - if on Windows use PowerShell Invoke-RestMethod instead"; \
	fi

clean:
	@echo "Removing containers and volumes (including DB data)..."
	$(COMPOSE) down -v --rmi local
