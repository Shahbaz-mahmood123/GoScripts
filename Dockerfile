# Use the official Golang image as the parent image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the source code to the working directory inside the container
COPY . .

# Build the application inside the container
RUN go build -o app

# Expose the port that the application will listen on
EXPOSE 8080

# Run the application inside the container
CMD ["./app"]
