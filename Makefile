SRC=src/main.go
BUILD_DIR=bin
BINARY=${BUILD_DIR}/auth

VERSION=1.0.0
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.DEFAULT_GOAL: clean build

build:
	go build -i ${LDFLAGS} -o ${BINARY} ${SRC}

# install:
# 	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

test:
	go test ./... -v

.PHONY: build install clean test
