FROM golang:1.16

WORKDIR /root
RUN go install github.com/shiv3/protoenv@v0.0.2

RUN mkdir -p .protoenv/protoc
RUN protoenv protoc init
RUN protoenv protoc install v3.17.3
RUN protoenv plugins protoc-gen-go install v1.26
