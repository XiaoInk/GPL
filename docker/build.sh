#!/bin/bash

export GOOS=linux
export CGO_ENABLED=0
export GO111MODULE=on
export GOPROXY="https://goproxy.io,direct"

[ ! -e 'bin' ] && mkdir bin

go build -a -installsuffix cgo -o bin/main cmd/main.go

# docker-compose -f docker/docker-compose-build.yml build
# docker push ...GPL:${VER:-latest}
