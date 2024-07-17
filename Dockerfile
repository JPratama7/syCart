FROM golang:1.22-alpine3.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]