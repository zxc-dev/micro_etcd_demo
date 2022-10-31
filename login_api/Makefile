BINARY_NAME=login_api

BIN_DIR=../bin

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -o ${BIN_DIR}/${BINARY_NAME}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BIN_DIR}/${BINARY_NAME} ] ; then rm ${BIN_DIR}/${BINARY_NAME} ; fi

