SRC=main.go
BUILD_DIR=bin
BINARY=${BUILD_DIR}/go-micro-auth

VERSION=1.0.0
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

default: test

build:
	go build -i ${LDFLAGS} -o ${BINARY} ${SRC}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

test:
	go test ./... -v

get-deps:
	go get github.com/dgrijalva/jwt-go
	go get github.com/golang/protobuf/proto
	go get github.com/hashicorp/consul/api
	go get github.com/micro/cli
	go get github.com/micro/go-micro
	go get github.com/micro/go-micro/client
	go get github.com/micro/go-micro/server
	go get golang.org/x/net/context
	go get github.com/stretchr/testify/require

.PHONY: build install clean test get-deps
