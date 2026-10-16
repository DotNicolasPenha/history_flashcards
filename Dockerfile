FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/main ./main

EXPOSE 8080

CMD ["./main"]