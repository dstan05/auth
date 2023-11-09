FROM golang:1.20.3-alpine AS builder

COPY .  /github.com/dstan05/auth/source/
WORKDIR /github.com/dstan05/auth/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/dstan05/auth/source/bin/auth_server .

CMD ["./auth_server"]