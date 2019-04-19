DIRECTORY_OUTPUT=build/_output/bin
FILE_OUTPUT=test

DIRECTORY_MAIN=./pkg/utility/

GOENV=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
GOFLAGS=-gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}"

TESTTARGETS := $(shell go list -e ./... | egrep -v "/(vendor)/")
# ex, -v
TESTOPTS := 


default: build

.PHONY: build
build:
	${GOENV} go build ${GOFLAGS} -o ${DIRECTORY_OUTPUT}/${FILE_OUTPUT} ${DIRECTORY_MAIN}

.PHONY: clean
clean:
	rm -rf ${DIRECTORY_OUTPUT}