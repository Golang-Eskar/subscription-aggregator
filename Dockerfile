FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем main.go прямо из cmd
RUN go build -o subscription-aggregator ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/subscription-aggregator .

# Копируем только .env
COPY .env ./

EXPOSE 8080

CMD ["./subscription-aggregator"]