FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app

COPY . /app

RUN go build -o stalcraftbot

FROM alpine:latest

WORKDIR /app

RUN  apk --update add \
        ca-certificates \
        && \
        update-ca-certificates

COPY --from=builder /app/stalcraftbot /app/stalcraftbot

COPY --from=builder /app/config.yaml /app/config.yaml

COPY --from=builder /app/00001_users_db.sql /app/00001_users_db.sql

ENTRYPOINT ["/app/stalcraftbot"]
