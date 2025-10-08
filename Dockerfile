FROM  golang:1.23.8-alpine3.21 as builder

RUN apk add --no-cache make ca-certificates gcc musl-dev linux-headers git jq bash

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app

COPY . /app/fishcake-service

WORKDIR /app/fishcake-service


RUN go mod download && go mod tidy
RUN make fishcake

FROM alpine:3.18

COPY --from=builder /app/fishcake /usr/local/bin
COPY --from=builder /app/fishcake-service/config.yaml /app/fishcake-service/config.yaml
COPY --from=builder /app/fishcakefrebase.json /app/fishcakefrebase.json
COPY --from=builder /app/migrations /app/migrations

ENV FISHCAKE_MIGRATIONS_DIR="/app/fishcake-service/migrations"
ENV FISHCAKE_CONFIG="/app/fishcake-service/config.yaml"

WORKDIR /app/fishcake-service
