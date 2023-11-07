# Dockerfile References: https://docs.docker.com/engine/reference/builder/


# Start from the latest golang base image
FROM golang:1.20

# Add Maintainer Info
#LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .
RUN go build -o main .
# Build the Go app
#RUN CGO_ENABLED=0 GOOS=linux go build -o main .


ENV PORT=3000
# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
