FROM golang:1.24-alpine

LABEL maintainer="https://github.com/C9b3rD3vi1" \
      version="1.0" \
      description="Go pastevault app service"

# Enable CGO for SQLite
ENV CGO_ENABLED=1 \
    GO111MODULE=on \
    GIN_MODE=release


# WORKING_DIR
WORKDIR /app

# Install git and other dependencies needed for go install
RUN apk add --no-cache git

# Install dependencies: git, gcc, musl-dev, sqlite-dev
RUN apk add --no-cache git gcc musl-dev sqlite-dev

# Install air for live reload
RUN go install github.com/air-verse/air@latest

# Pre-copy go.mod and go.sum to use Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container app
COPY . .

EXPOSE 3000

# Start using air
CMD ["air", "-c", ".air.toml"]

# Optional: for production builds, remove air and use this instead
# RUN go build -o main main.go
# CMD ["./main"]


# termina build
#docker build -t pastevault-dev .
# terminal run
# docker run -it --rm -p 3000:3000 -v $(pwd):/app pastevault-dev
