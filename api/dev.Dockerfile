FROM golang:1.25.5-alpine3.23

RUN apk update && apk add --no-cache curl

RUN go install github.com/air-verse/air@latest

WORKDIR /app

CMD ["/bin/sh", "-c", "go mod download && air -c .air.toml"]
