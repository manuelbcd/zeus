ARG GO_VERSION=latest
FROM golang:${GO_VERSION}

ARG PROJECT_NAME=project
ARG GO_PORT=8080
ARG GOPATH=/go
ARG PROJECT_CONTAINER_PATH=/go/src

RUN apt-get update && \
    apt-get install git

RUN go get github.com/codegangsta/gin \
           github.com/gorilla/mux \
           gopkg.in/mgo.v2 \
           golang.org/x/oauth2 \
           github.com/google/uuid \
           github.com/google/go-github/github \
           github.com/garyburd/redigo/redis

RUN mkdir -p ${PROJECT_CONTAINER_PATH}
WORKDIR ${PROJECT_CONTAINER_PATH}

ENV BIN_APP_PORT ${GO_PORT}
ENV GOPATH ${GOPATH}
EXPOSE ${GO_PORT}

CMD ["gin", "run", "server.go"]
