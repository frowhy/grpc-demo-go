FROM golang:latest

WORKDIR $GOPATH/src/test

ADD . $GOPATH/src/test

RUN go build

EXPOSE 50051

ENTRYPOINT  ["./test"]