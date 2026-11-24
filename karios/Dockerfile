# Stufe 1: Go-Webserver statisch für x86_64 kompilieren
FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY main.go .
RUN GOOS=linux GOARCH=amd64 go build -o /webserver main.go

# Stufe 2: Fertiges Hadron-Kairos-Core-Image (x86_64) + dein Webserver
FROM quay.io/kairos/hadron:v0.6.3-core-amd64-generic-v4.7.0

COPY --from=builder /webserver /usr/bin/webserver
COPY public /var/www/public
