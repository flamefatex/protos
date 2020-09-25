FROM golang:1.14-alpine3.12

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache git wget protobuf bash

RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-linux-x86_64.zip
RUN unzip protoc-3.12.3-linux-x86_64.zip -d /usr/local/ -x bin/protoc
RUN rm protoc-3.12.3-linux-x86_64.zip

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
RUN go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
RUN go get -u github.com/golang/mock/gomock@v1.4.4
RUN go get -u github.com/golang/mock/mockgen@v1.4.4

RUN git clone https://github.com/grpc-ecosystem/grpc-gateway.git /go/src/github.com/grpc-ecosystem/grpc-gateway
RUN git clone https://github.com/mwitkow/go-proto-validators.git /go/src/github.com/mwitkow/go-proto-validators
RUN go get golang.org/x/net@v0.0.0-20191002035440-2ec189313ef0

WORKDIR /opt/protos