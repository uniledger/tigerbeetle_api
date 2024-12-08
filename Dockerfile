FROM golang:1.22-alpine AS builder

RUN apk update && apk add --no-cache git build-base zig

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=1 CC="zig cc"

RUN go build -o tigerbeetle_api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/tigerbeetle_api .
ENTRYPOINT ["./tigerbeetle_api"]
