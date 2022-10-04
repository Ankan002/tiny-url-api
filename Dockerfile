FROM golang:alpine AS builder

WORKDIR /usr/tiny-url-api

ARG GO_ENV
ARG PORT
ARG DB_URI

COPY go.mod .
COPY go.sum .

RUN ["go", "mod", "download"]

COPY . .

RUN ["go", "build", "-o", "/build"]

CMD ["/build"]
