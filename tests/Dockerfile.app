ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS builder

ARG ID_TYPE

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
COPY . /firelete
WORKDIR /firelete/tests/integration-${ID_TYPE}-id
RUN go mod download
RUN go build -o ./app ./main.go

FROM alpine:latest

ARG ID_TYPE

WORKDIR /root/
COPY --from=builder ./firelete/tests/integration-${ID_TYPE}-id/app ./
EXPOSE 8080
ENTRYPOINT ["./app"]