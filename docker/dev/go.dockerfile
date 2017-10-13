ARG GO_VERSION=latest
FROM golang:${GO_VERSION}

ARG PROJECT_NAME=project
ARG GO_PORT=8080

RUN go get github.com/codegangsta/gin \
		   github.com/gorilla/mux \
           gopkg.in/mgo.v2

RUN mkdir /go/src/${PROJECT_NAME}
WORKDIR /go/src/${PROJECT_NAME}

ENV BIN_APP_PORT ${GO_PORT}
EXPOSE ${GO_PORT}

CMD ["gin", "run", "main.go"]
