FROM golang:1.25.6-trixie AS builder

WORKDIR /go/src/flare
COPY . .

RUN apt update && \
    apt install -y libsystemd-dev && \
    go build -o /go/bin/flare

FROM debian:trixie

WORKDIR /app
COPY --from=builder /go/bin/flare /app/flare
ENTRYPOINT [ "/app/flare", "run"]
