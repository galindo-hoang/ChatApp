FROM golang:1.21.2

# make directory is root folder in images
WORKDIR /app
# copy go-mod and go-sum to current root directory
COPY go.mod go.sum ./

RUN go mod download


# This COPY command uses a wildcard to copy all files with .go extension located in the current directory
# on the host (the directory where the Dockerfile is LOCATED) into the current directory inside the image.
COPY ./ ./

# The result of that command will be a static application binary named docker-gs-ping
# and located in the root of the filesystem of the IMAGE that you are building.
RUN CGO_ENABLED=0 GOOS=linux go build -o /chatapp ./cmd/chatapp

CMD ["/chatapp"]

