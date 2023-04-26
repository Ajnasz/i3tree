BINARY_NAME=i3tree
VERSION=$(shell git describe --tags)
BUILD=$(shell date +%FT%T%z)
BUILD_ARGS=-trimpath -ldflags '-w -s'

.PHONY: build clean linux freebsd

all: clean linux freebsd

${BINARY_NAME}.linux.amd64:
	GOARCH=amd64 GOOS=linux go build ${BUILD_ARGS} -o $@ main.go

${BINARY_NAME}.linux.arm64:
	GOARCH=arm64 GOOS=linux go build ${BUILD_ARGS} -o $@ main.go

${BINARY_NAME}.linux.386:
	GOARCH=386 GOOS=linux go build ${BUILD_ARGS} -o $@ main.go

${BINARY_NAME}.freebsd.amd64:
	GOARCH=amd64 GOOS=freebsd go build ${BUILD_ARGS} -o $@ main.go

${BINARY_NAME}.freebsd.arm64:
	GOARCH=arm64 GOOS=freebsd go build ${BUILD_ARGS} -o $@ main.go

${BINARY_NAME}.freebsd.386:
	GOARCH=386 GOOS=freebsd go build ${BUILD_ARGS} -o $@ main.go

linux: ${BINARY_NAME}.linux.amd64 ${BINARY_NAME}.linux.arm64 ${BINARY_NAME}.linux.386
freebsd: ${BINARY_NAME}.freebsd.amd64 ${BINARY_NAME}.freebsd.arm64 ${BINARY_NAME}.freebsd.386

clean:
	rm -f ${BINARY_NAME}.*.*
