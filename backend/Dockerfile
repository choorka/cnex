# Этап сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Сборка бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app

# Финальный образ с сертификатами
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/app /

# Запуск как non-root пользователь
USER nonroot
ENTRYPOINT ["/app"]
