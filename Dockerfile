FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

# Install wget and download migrate as .tar.gz
RUN apt-get update && apt-get install -y wget && \
    ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then ARCH="amd64"; fi && \
    if [ "$ARCH" = "aarch64" ]; then ARCH="arm64"; fi && \
    wget -O /tmp/migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-${ARCH}.tar.gz && \
    tar -xzf /tmp/migrate.tar.gz -C /usr/local/bin && \
    chmod +x /usr/local/bin/migrate && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

COPY . .

RUN go build -o bot .

CMD ["./bot"]
