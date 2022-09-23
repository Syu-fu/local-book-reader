FROM golang:1.19.1-alpine

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go install github.com/cespare/reflex@latest
ENV CGO_ENABLED=0

WORKDIR /go/src/api
COPY ./golang ./

RUN go mod download
