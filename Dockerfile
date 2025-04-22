FROM  golang:1.23.8-alpine3.21 as builder

RUN apk add --no-cache make ca-certificates gcc musl-dev linux-headers git jq bash

WORKDIR /app

COPY . /app

RUN go mod download && go mod tidy
RUN make fishcake

FROM alpine:3.18

COPY --from=builder /app/fishcake /usr/local/bin
COPY --from=builder /app/config.yaml /app/config.yaml
COPY --from=builder /app/fishcakefrebase.json /app/fishcakefrebase.json
COPY --from=builder /app/migrations /app/migrations

WORKDIR /app
#ENV INDEXER_MIGRATIONS_DIR="/app/acorus/migrations"
