# Makefile для создания миграций
include .env
# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATE := migrate -path ./migrations -database "$(DB_DSN)"

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go # Теперь при вызове make run мы запустим наш сервер

gen:
	oapi-codegen -config ./api/openapi-config.yaml ./api/openapi.yaml > ./internal/handlers/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number