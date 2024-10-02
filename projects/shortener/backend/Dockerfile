# Use the latest Go Alpine image
FROM golang:1.22.5-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN cd cmd && go build -o main .

# Expose the port your Go application listens on
EXPOSE 3000

# Command to run the Go application and load environment variables from .env file
CMD ["sh", "-c", "cd cmd && ./main"]