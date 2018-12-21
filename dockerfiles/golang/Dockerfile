FROM golang:1.11.3-stretch

RUN apt-get update && \
    apt-get install -y \
        default-libmysqlclient-dev \
        mysql-client \
        vim \
        autoconf \
        automake \
        libtool \
        curl \
        make \
        g++ \
        unzip

ARG SERVER_APP_PATH="/go/src/github.com/endotakuya/grpc-example"
WORKDIR $SERVER_APP_PATH

# install protof
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip && \
    unzip protoc-3.6.1-linux-x86_64.zip -d protoc3  && \
    mv protoc3/bin/* /usr/local/bin/  && \
    mv protoc3/include/**/* /usr/local/include/**/*

COPY . .
RUN go get -u \
        github.com/golang/dep/cmd/dep \
        google.golang.org/grpc \
        github.com/golang/protobuf/protoc-gen-go
