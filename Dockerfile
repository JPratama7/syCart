FROM docker.io/golang:1.22-alpine3.20 as builder

WORKDIR /app

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main cmd/api/main.go

FROM docker.io/alpine:3.20

RUN apk add --no-cache ca-certificates dumb-init

COPY --from=builder /app/main /app/main

ENTRYPOINT ["dumb-init --"]

CMD ["/app/main"]