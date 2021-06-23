FROM golang:1.15.2-alpine

RUN apk update && apk add git

RUN mkdir /go/src/app

WORKDIR /go/src/app
