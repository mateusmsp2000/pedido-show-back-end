# Dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN apk update

RUN apk add \
  g++ \
  git \
  musl-dev \
  go \
  tesseract-ocr-dev

COPY . .

RUN CGO_ENABLED=1 go build -ldflags "-s -w" -a -o main

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
