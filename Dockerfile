FROM golang:alpine

ENV CGO_ENABLED=0

RUN apk update && apk add git && go get -u golang.org/x/lint/golint \
    && go get github.com/stretchr/testify

ADD ./go.mod /go/src/go.mod
ADD ./go.sum /go/src/go.sum
ADD ./main.go /go/src/main.go
ADD ./counter /go/src/counter

WORKDIR /go/src

RUN go build -o main .
