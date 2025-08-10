
FROM golang:1.24 AS builder
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o state-machine-telegram-bot .

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/state-machine-telegram-bot /state-machine-telegram-bot
EXPOSE 8000
CMD ["/state-machine-telegram-bot", "runbot"]