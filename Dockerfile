# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o opa-control-plane 

# Final stage
FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y curl bash && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/opa-control-plane .
ENTRYPOINT ["/app/opa-control-plane", "run"]
