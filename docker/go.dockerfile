ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION}

ARG PROJECT_NAME=project
RUN go get github.com/gorilla/mux gopkg.in/mgo.v2 github.com/codegangsta/gin

RUN mkdir /go/src/${PROJECT_NAME}si
WORKDIR /go/src/${PROJECT_NAME}

EXPOSE 8080

CMD ["go", "run", "main.go"]
