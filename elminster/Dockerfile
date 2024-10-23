FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o dist/elminster

# ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dist/elminster .

CMD ["./elminster", "server"]

