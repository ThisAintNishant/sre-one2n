# ---------- Builder ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o server \
    ./cmd/server

# ---------- Runtime ----------
FROM alpine:3.22

RUN adduser -D appuser

WORKDIR /app

COPY --from=builder /app/server .

USER appuser

EXPOSE 8080

CMD ["./server"]