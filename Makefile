DATE := $(shell TZ=Asia/Tokyo date +%Y%m%d-%H:%M:%S)
HASH := $(shell git rev-parse --short HEAD)
GOVERSION := $(shell go version)
LDFLAGS := -X 'main.date=${DATE}' -X 'main.hash=${HASH}' -X 'main.goversion=${GOVERSION}'

NAME := aie-emu
MODULE := github.com/sfuruya0612/${NAME}

.PHONY: build

lint:
	golangci-lint run

build: lint
	-rm -rf build
	mkdir build

	go mod tidy

	GOOS=linux GOARGH=amd64 go build -ldflags "${LDFLAGS}" ${MODULE}
	zip build/${NAME}_linux_amd64.zip ${NAME}

	GOOS=darwin GOARGH=amd64 go build -ldflags "${LDFLAGS}" ${MODULE}
	zip build/${NAME}_darwin_amd64.zip ${NAME}

	@rm ${NAME}

install: lint
	-rm ${GOPATH}/bin/${NAME}
	go mod tidy
	go install -ldflags "${LDFLAGS}" ${MODULE}

clean:
	-rm ${GOPATH}/bin/${NAME}
	-rm -rf build
