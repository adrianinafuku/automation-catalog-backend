# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# Run stage
FROM alpine:3.19

WORKDIR /root/

COPY --from=builder /app/server .

ENV PORT=8080
EXPOSE 8080

CMD ["./server"]
