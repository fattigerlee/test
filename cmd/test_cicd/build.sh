#!/bin/bash

echo "编译..."

GOPROXY="https://goproxy.cn" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

echo "编译完成..."

echo "拷贝配置文件..."

echo "拷贝配置文件完成..."
