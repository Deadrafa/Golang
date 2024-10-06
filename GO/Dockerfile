# Шаг 1: Используем официальный образ Go для сборки
FROM golang:1.21-alpine AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum (если у вас есть)
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Сборка Go-приложения
RUN go build -o shavel

# Шаг 2: Создаем финальный образ с минимальным размером
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный бинарный файл из этапа сборки
COPY --from=build /app/shavel .

# Копируем папки static и templates
COPY --from=build /app/static ./static
COPY --from=build /app/templates ./templates

# Открываем порт для HTTP-сервера
EXPOSE 8080

# Команда для запуска приложения
CMD ["./shavel"]

