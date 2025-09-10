# Build stage
FROM golang:1.24.0 AS builder
WORKDIR /app
# Copy go.mod and go.sum first, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o opa-control-plane 

# Final stage
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y curl bash && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/opa-control-plane .
ENTRYPOINT ["/app/opa-control-plane"]
CMD ["run"]
