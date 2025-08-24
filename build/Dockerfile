# Простой одноэтапный билд
FROM golang:1.24-alpine

WORKDIR /app

# Копируем всё
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/app

COPY migrations /app/migrations

# Запускаем
CMD ["./main"]