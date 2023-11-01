# Use the official Go image as the base image
FROM golang:1.20 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o restAPI cmd/restAPI/restAPI.go

# Use a minimal Alpine Linux image as the final base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the build image
COPY --from=build /app/restAPI .

# Expose the port that the application listens on
EXPOSE 8080

# Define the command to run the application
CMD ["./restAPI"]
