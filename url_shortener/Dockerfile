FROM golang:1.23-alpine

WORKDIR /app

# Копируем исходники проекта
COPY . .

# Объявляем переменную окружения, указывающую путь к конфигу внутри контейнера
ENV CONFIG_PATH="/app/config/config.yaml"

RUN go mod tidy && go build -o app ./cmd/app

EXPOSE 8080

CMD ["./app"]