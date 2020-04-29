### Stage 1
FROM golang:1.13-alpine3.10 AS builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/app -v

#### Stage 2
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/build/app .
COPY --from=builder /app/input.json .
ENV SOURCE_FILE "input.json"
CMD ["./app"]
