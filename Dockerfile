FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8081

CMD ["./server"]