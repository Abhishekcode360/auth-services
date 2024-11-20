FROM golang:1.22.8 AS builder

WORKDIR /var/www/auth-api/src

COPY src .

EXPOSE 3001

ENTRYPOINT [ "/bin/bash","-c","go run main.go" ]