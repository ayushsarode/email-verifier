# Use Golang as a builder
FROM golang:1.22 AS builder

WORKDIR /app

# Cache dependencies to speed up builds
COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:latest

WORKDIR /root/


COPY --from=builder /app/app .

RUN chmod +x ./app

# Run the CLI tool by default
CMD ["./app"]
