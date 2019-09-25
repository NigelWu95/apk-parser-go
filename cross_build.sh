#!/usr/bin/env bash
DIR=$(cd ../; pwd)
export GOPATH=${GOPATH}:${DIR}
GOOS=linux GOARCH=amd64 go build -o ../bin/apk-parser-linux_amd64
cp ../bin/apk-parser-linux_amd64 ../deploy/
