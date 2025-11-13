FROM golang:1.22-alpine AS builder


RUN apk add --no-cache git build-base


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN go build -o subscription-aggregator ./cmd/app

FROM alpine:latest


RUN apk --no-cache add ca-certificates tzdata


WORKDIR /app


COPY --from=builder /app/subscription-aggregator .


COPY .env config.yaml ./

EXPOSE 8080


CMD ["./subscription-aggregator"]