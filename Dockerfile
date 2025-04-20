# Используем официальный образ Go
FROM golang:1.24.2

# Устанавливаем Air
RUN go install github.com/air-verse/air@latest

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Сборка приложения
RUN go build -o /app/tmp/main ./cmd

# Запуск Air для горячей перезагрузки
CMD ["air"]
