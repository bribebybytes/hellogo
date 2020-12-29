FROM golang:latest

WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go get -d github.com/gorilla/mux

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]