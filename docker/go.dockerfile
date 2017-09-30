ARG GO_VERSION=latest
FROM golang:${GO_VERSION}

ARG PROJECT_NAME=project
ARG GO_PORT=80

RUN go get github.com/gorilla/mux \
           gopkg.in/mgo.v2 \
	       github.com/codegangsta/gin

RUN mkdir /go/src/${PROJECT_NAME}
WORKDIR /go/src/${PROJECT_NAME}

ENV LISTENING_ADDRESS=:${GO_PORT}
ENV BIN_APP_PORT=${GO_PORT}

EXPOSE ${GO_PORT}

CMD ["gin", "run", "main.go"]
