# syntax=docker/dockerfile:1

FROM golang:1.24 AS build

WORKDIR $GOPATH/src/github.com/brotherlogic/tasklister

COPY go.mod ./
COPY go.sum ./

RUN mkdir server
COPY server/*.go ./server/



RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /tasklister

##
## Deploy
##
FROM ubuntu:22.04
USER root:root

WORKDIR /
COPY --from=build /tasklister /tasklister
RUN echo "" > /known_hosts

ENTRYPOINT ["/tasklister"]