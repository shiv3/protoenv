FROM golang:1.16

WORKDIR /
RUN go install github.com/shiv3/protoenv@latest

RUN protoenv protoc install v3.17.3
RUN protoenv plugins protoc-gen-go install v1.20.0
