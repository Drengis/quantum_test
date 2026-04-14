FROM golang:1.26-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/quantum-server ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/.env .
COPY --from=builder /app/quantum-server .
COPY --from=builder /app/docs ./docs

EXPOSE 8081

CMD ["./quantum-server"]