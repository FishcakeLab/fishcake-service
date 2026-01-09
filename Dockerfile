FROM  golang:1.24.3-alpine3.21 as builder

RUN apk add --no-cache make ca-certificates gcc musl-dev linux-headers git jq bash

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app

RUN go mod download

COPY . /app/fishcake-service

WORKDIR /app/fishcake-service

RUN make fishcake

FROM alpine:3.18

COPY --from=builder /app/fishcake-service/fishcake /usr/local/bin
COPY --from=builder /app/fishcake-service/config.yaml /app/fishcake-service/config.yaml
COPY --from=builder /app/fishcake-service/fishcakefrebase.json /app/fishcake-service/fishcakefrebase.json
COPY --from=builder /app/fishcake-service/migrations /app/fishcake-service/migrations

ENV FISHCAKE_FFB="/app/fishcake-service/fishcakefrebase.json"
ENV FISHCAKE_MIGRATIONS_DIR="/app/fishcake-service/migrations"
ENV FISHCAKE_CONFIG="/app/fishcake-service/config.yaml"

WORKDIR /app/fishcake-service
