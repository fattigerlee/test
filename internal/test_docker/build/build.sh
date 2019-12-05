#!/bin/bash

echo "编译程序..."

CGO_ENABLED=0
GOOS=linux
GOARCH=amd64

#工作路径
workdir=$(pwd)

#程序名称
name=test

go build -o=$workdir/$name $workdir/../main.go
#go build $workdir/../main.go

echo "完成"

echo "拷贝资源..."

echo "完成"

echo "生成Dockfile..."

cat>Dockerfile<<EOF
FROM centos

WORKDIR /root/

COPY test ./app/

CMD ["./app/test"]
EOF

echo "完成"

echo "构建镜像..."

docker build -t $name .

echo "完成"


