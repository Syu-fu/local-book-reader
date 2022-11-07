FROM golang:1.19.1-alpine

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go install github.com/cespare/reflex@latest
ENV CGO_ENABLED=0

WORKDIR /go/src/api
COPY ./golang ./
COPY ./.env ./.env

RUN go mod download
