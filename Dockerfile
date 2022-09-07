# goバージョン
FROM golang:1.16.3-alpine
RUN apk add --update &&  apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app