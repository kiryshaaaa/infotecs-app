# Используем официальный образ Go
FROM golang:alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей и зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/infotecs-app/main.go

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]