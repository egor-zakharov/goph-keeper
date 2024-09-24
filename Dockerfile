# Используем официальный образ Go для сборки (этап сборки)
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные исходные файлы
COPY . .

# Собираем бинарный файл приложения
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server/main.go

# Используем минимальный базовый образ для запуска приложения (этап запуска)
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный бинарник из предыдущего этапа
COPY --from=builder /app/server .

# Если есть дополнительные файлы конфигурации, скопируйте их
# COPY --from=builder /app/config.yaml .

# Указываем, какой порт будет использовать приложение (если необходимо)
EXPOSE 8081

# Задаем команду для запуска приложения
CMD ["./server"]