FROM golang:1.17 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY cmd cmd
COPY config config
COPY internal internal
COPY pkg pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/fibonacci_service

FROM alpine:latest
LABEL maintainer="Abdulla Akhmedov <abdullapower@outlook.com>"
LABEL org.opencontainers.image.source=https://github.com/Eretic431/fibonacci
WORKDIR /app
COPY --from=builder /build/app .

CMD ["./app"]