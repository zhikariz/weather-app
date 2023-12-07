# Use a More Specific Golang Base Image
FROM golang:1.21.5-alpine3.18

# Set the time zone
RUN apk update && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone && \
    apk del tzdata

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod .
COPY go.sum .

# Download and install dependencies
RUN export GOPROXY=https://proxy.golang.org && \
    go mod tidy

# Copy the entire application to the container
COPY . .

# Build the Golang application
RUN go build -o main cmd/server/main.go

# Remove unnecessary files after the build
RUN rm -rf go.mod go.sum

# Create a non-root user for running the application
RUN adduser -D -g '' myuser
USER myuser

# Expose the port used by the application
EXPOSE 8080

# Command to run the application
CMD ["./main"]