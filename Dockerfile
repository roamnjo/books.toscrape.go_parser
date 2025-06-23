FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o parser ./cmd/main.go

FROM debian:bookworm-slim

COPY --from=builder /app/parser /usr/local/bin/parser

WORKDIR /usr/local/bin

CMD ["parser"]