# Use the official Golang image as a base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o myapp .

# Command to run the executable
CMD ["./myapp"]

# FROM golang:1.18-alpine

# WORKDIR /app

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# COPY . .

# RUN go build -o /myapp

# EXPOSE 8080

# CMD ["/myapp"]
