############################
# STEP 1 build executable binary
############################

# Based off the Dockerfile example at https://hub.docker.com/_/golang

# Use the official Golang runtime as a parent image
FROM golang:latest AS builder

# Set the current working directory
# WORKDIR /go/src/app
WORKDIR /app/backend

# Copy the current directory contents into the container at /app
COPY . .

# Download and install necessary Golang dependencies
# RUN go get -d -v ./...
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v ./...

# Make port 8080 available to the world outside this container
# EXPOSE 8080

# Run the executable when the container launches
# CMD ["api"]

######## Start a new stage from scratch #######
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
# COPY --from=builder /go/src/app/api .
COPY --from=builder /main ./

RUN chmod +x ./main
ENTRYPOINT ["./main"]

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
# CMD ["./api"]
