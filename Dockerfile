# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.21.6-alpine

# Install make
RUN apk add --no-cache make

# Environment variables which CompileDaemon requires to run
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# Basic setup of the container
RUN mkdir /app
COPY .. /app
WORKDIR /app

# Add go path on .bashrc
RUN echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bashrc
RUN echo "export PATH=$PATH:$(go env GOROOT)/bin" >> ~/.bashrc

# Get CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

# Get Swag
RUN go get -u github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

# The build flag sets how to build after a change has been detected in the source code
# The command flag sets how to run the app after it has been built only
ENTRYPOINT CompileDaemon -build="make build" -command="make" -color=true -log-prefix=false -exclude-dir=docs
