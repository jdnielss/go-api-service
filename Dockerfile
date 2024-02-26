# Use the official golang image as a base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
