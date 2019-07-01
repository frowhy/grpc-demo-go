FROM golang:latest

WORKDIR /grpc-demo-go

ADD . /grpc-demo-go

RUN export GO111MODULE=on \
    && export GOPROXY=https://goproxy.io \
    && go build

EXPOSE 50051

ENTRYPOINT  ["./grpc-demo-go"]