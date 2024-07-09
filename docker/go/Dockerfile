FROM golang:1.22-alpine

ARG app_env
ENV APP_ENV $app_env

RUN apk add --no-cache \
    bash \
    git \
    openssh-client

RUN mkdir /app
WORKDIR /app
COPY . /app

# Build application
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build