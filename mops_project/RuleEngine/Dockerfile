# Используем образ Go (версию подставьте свою)
FROM golang:1.20-alpine AS builder

# Создаём рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum в /app
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы (включая папку cmd/)
COPY . .

# Собираем бинарник в /app/bin
RUN go build -o /app/bin/engine ./main.go

# Минимизируем размер образа, используя stage multistage
FROM alpine:3.17
WORKDIR /app

# Копируем бинарник из первого stage
COPY --from=builder /app/bin/engine .

# Запускаем бинарь
CMD ["./engine"]
