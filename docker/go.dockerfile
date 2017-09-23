ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION}

ARG PROJECT_NAME=project
RUN go get github.com/gorilla/mux gopkg.in/mgo.v2

RUN mkdir /go/src/${PROJECT_NAME}
WORKDIR /go/src/${PROJECT_NAME}

ENV BASE_URL=localhost
ENV PORT=8080

EXPOSE 8080
CMD ["go", "run", "main.go"]
