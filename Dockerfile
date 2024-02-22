FROM --platform=$BUILDPLATFORM golang:1.21.1-alpine3.18 as builder

RUN apk add --no-cache make ca-certificates gcc musl-dev linux-headers git jq bash

WORKDIR /app

COPY . /app

RUN go mod download && go mod tidy
RUN make acorus

FROM alpine:3.18

COPY --from=builder /app/fishcake /usr/local/bin
COPY --from=builder /app/config.yaml /app/config.yaml
COPY --from=builder /app/migrations /app/migrations

WORKDIR /app
#ENV INDEXER_MIGRATIONS_DIR="/app/acorus/migrations"
