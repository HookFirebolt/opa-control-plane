FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y curl bash && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY opa-control-plane /app/opa-control-plane
ENTRYPOINT ["/app/opa-control-plane"]
CMD ["run"]