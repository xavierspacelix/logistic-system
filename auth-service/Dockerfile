# syntax=docker/dockerfile:1
FROM golang:1.24.5-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-service .

EXPOSE 8001

CMD ["./auth-service"]
