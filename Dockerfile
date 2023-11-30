# Builder stage
FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/go-task/task/v3/cmd/task@latest

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN task build

# For minimal image
FROM scratch

COPY --from=builder /build/http-server /http-server
COPY --from=builder /build/.env .env

ENV GO_ENV=production

CMD ["./http-server"]