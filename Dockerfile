FROM golang:1.22.5 AS builder

COPY . /

WORKDIR /

RUN go build ./main.go

FROM alpine:latest

COPY --from=builder /go/main .
