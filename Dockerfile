FROM golang:1.24 AS builder

WORKDIR /app

# copiar mod y sum
COPY go.mod go.sum ./
RUN go mod download

# copiar todo el repo
COPY . .

# entrar al paquete real
WORKDIR /app/cmd/server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server .

FROM alpine:3.19
WORKDIR /root/
COPY --from=builder /server .
ENV PORT=8080
EXPOSE 8080
CMD ["./server"]
