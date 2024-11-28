FROM golang:latest AS builder

RUN apt install git

WORKDIR /app
COPY . .
RUN go build -o tigerbeetle_api .

# Final stage - debian:slim has the libc we need
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/tigerbeetle_api .

ENTRYPOINT ["./tigerbeetle_api"]
