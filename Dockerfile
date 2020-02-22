ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /go/src/app/app .

EXPOSE 8080

CMD ["/app/app"]