FROM golang:1.22.2

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the container
COPY . ./

EXPOSE 8080

# running makefile
CMD ["make", "run","make","migrateUP"]