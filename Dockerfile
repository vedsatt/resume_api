# образ Go на базе легковесной системы alpine
FROM golang:1.23.7-alpine3.21 AS server

# переключаемся на рабочую директорию
WORKDIR /server

# копируем модуль программы
COPY go.mod ./

# обновляем зависимости
RUN go mod tidy

# копируем все файлы программы в образ
COPY . .

# собираем программу 
RUN CGO_ENABLED=0 go build -o main ./cmd

# стандартный порт для развертывания сервера
EXPOSE 8080


