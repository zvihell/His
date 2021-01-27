FROM golang:latest

WORKDIR /backend

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Export necessary port
EXPOSE 8888


# Command to run when starting the container
CMD ["./main"]