# build stage
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server-app ./cmd

# final stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server-app ./

COPY --from=builder /app/database/migrations ./database/migrations

COPY .env .

EXPOSE 8080

CMD ["./server-app"]



