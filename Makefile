include .env
export


export PROJECT_ROOT=$(shell pwd)

env-up: 
	@docker compose up -d bt-postgres

env-down: 
	@docker compose down bt-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		make env-down && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi;

env-port-forward:
	docker compose up -d port-forwarder

env-port-close:
	docker compose down port-forwarder

ps: 
	@docker compose ps

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутвует обязательный параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm bt-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)" \

migrate-action: 
	@if [ -z "$(action)" ]; then \
		echo "Отсутвует обязательный параметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm bt-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@bt-postgres:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable \
		"$(action)"
migrate-up:
	@make migrate-action action=up
migrate-down:
	@make migrate-action action=down
	