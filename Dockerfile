# Define go base image and version
FROM golang:1.20.3-alpine

# Install g++ dependency (for plugin build)
RUN apk --update add g++

# Define GOROOT variable
ENV GOROOT=/usr/local/go

# Change base path
WORKDIR /go/code/src