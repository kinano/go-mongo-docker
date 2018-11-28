# This is a generic golang container that can be used to build other images
# Accept the Go version for the image to be set as a build argument
# Default to 1.11
ARG GOLANG_VERSION=1.11

FROM golang:${GOLANG_VERSION}

# Install hot reloader
RUN go get github.com/githubnemo/CompileDaemon