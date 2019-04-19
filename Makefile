DIRECTORY_OUTPUT=build/_output/bin
FILE_OUTPUT=test

DIRECTORY_MAIN=./cmd/service/

GOENV=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
GOFLAGS=-gcflags="all=-trimpath=${GOPATH}" -asmflags="all=-trimpath=${GOPATH}"

TESTTARGETS := $(shell go list -e ./... | egrep -v "/(vendor)/")
# ex, -v
TESTOPTS := 


default: build

.PHONY: format
format:
	go fmt $(shell find -type f -name '*.go' | grep -v vendor | sed 's|\(.*\)/[^/]*|\1|g')

${DIRECTORY_OUTPUT}/${FILE_OUTPUT}:
	${GOENV} go build ${GOFLAGS} -o ${DIRECTORY_OUTPUT}/${FILE_OUTPUT} ${DIRECTORY_MAIN}

.PHONY: build
build: ${DIRECTORY_OUTPUT}/${FILE_OUTPUT}

.PHONY: clean
clean:
	rm -rf ${DIRECTORY_OUTPUT}

.PHONY: run
run: build
	${DIRECTORY_OUTPUT}/${FILE_OUTPUT}