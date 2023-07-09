# Используем официальный образ Golang как базовый
FROM golang:latest as builder

# Установим рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go mod и sum файлы
COPY go.mod go.sum ./

# Скачиваем все зависимости
RUN go mod download

# Копируем исходный код в рабочую директорию контейнера
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o lcrm2 .

# Начинаем новую стадию сборки для создания минимального образа
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем исполняемый файл из первой стадии во вторую
COPY --from=builder /app/lcrm2 .
COPY --from=builder /app/.env .
EXPOSE 8080
# Запускаем приложение
CMD ["./lcrm2"]