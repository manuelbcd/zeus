ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION}

ARG PROJECT_NAME=project
ARG GO_PORT=80


RUN go get github.com/gorilla/mux \
	gopkg.in/mgo.v2 \
	github.com/codegangsta/gin

RUN mkdir /go/src/${PROJECT_NAME}
WORKDIR /go/src/${PROJECT_NAME}

EXPOSE ${GO_PORT}

CMD ["go", "run", "main.go"]
