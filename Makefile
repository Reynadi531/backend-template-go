BUILD_DIR = bin
VERSION := $(shell git describe --tags --always --dirty)
BUILD := $(shell date +%Y-%m-%d\ %H:%M)
LDFLAGS=-ldflags="-w -s -X 'libcommon.Version=${VERSION}' -X 'libcommon.Build=${BUILD}'"

.PHONY: clean

clean:
	rm -rf ./tmp && rm -rf ./bin

build:
	go build ${LDFLAGS} -o $(BUILD_DIR)/ cmd/main.go

docker.build:
	docker build -f docker/Dockerfile -t backend-template-go:latest .

docker.build.alpine:
	docker build -f docker/alpine.Dockerfile -t backend-template-go:latest .

docker.dev:
	docker-compose -f docker/docker-compose.dev.yaml up

docker.dev.build:
	docker-compose -f docker/docker-compose.dev.yaml up --build