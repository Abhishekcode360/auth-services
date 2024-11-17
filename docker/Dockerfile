FROM golang:1.22.8 AS builder

WORKDIR /var/www/auth-api/src

COPY src .

RUN go build -o main

ENTRYPOINT [ "/bin/bash","-c","./main" ]