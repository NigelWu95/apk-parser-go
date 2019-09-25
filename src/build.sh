#!/usr/bin/env bash
DIR=$(cd ../; pwd)
export GOPATH=$GOPATH:$DIR
go build -o ../bin/apk-parser
