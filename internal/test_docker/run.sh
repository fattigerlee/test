#!/bin/bash

echo "编译..."

go build .
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .