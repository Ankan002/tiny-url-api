FROM golang:alpine AS builder

ARG GO_ENV
ARG PORT

WORKDIR /usr/tiny-url-api

COPY go.mod .
COPY go.sum .

RUN ["go", "mod", "download"]

COPY . .

RUN ["go", "build", "-o", "/build"]

CMD ["/build"]
