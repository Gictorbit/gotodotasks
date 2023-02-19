# syntax=docker/dockerfile:1.4
ARG GOOSE_VERSION="v3.9.0"
FROM alpine:latest as final
# Install goose binary
RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ARG GOOSE_VERSION
ADD "https://github.com/pressly/goose/releases/download/${GOOSE_VERSION}/goose_linux_x86_64" /bin/goose
RUN chmod +x /bin/goose

WORKDIR /migrations
COPY ./sqls ./sqls
RUN ls -ltrh
ENTRYPOINT ["goose"]
