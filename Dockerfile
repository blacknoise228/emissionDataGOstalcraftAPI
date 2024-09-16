FROM golang:latest AS builder

WORKDIR /app

COPY . /app

RUN go build -o stalcraftbot

FROM debian:latest

COPY --from=builder /app/stalcraftbot /app/stalcraftbot

COPY --from=builder /app/config.yaml /app/config.yaml

ENTRYPOINT ["/app/stalcraftbot"]
