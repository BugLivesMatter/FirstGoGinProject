# Используем минимальный образ Go
FROM golang:1.25-alpine

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем go-модули и зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем бинарник
RUN go build -o app

# Открываем порт
EXPOSE 4200

# Запуск приложения
CMD ["./app"]