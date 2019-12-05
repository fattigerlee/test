#!/bin/sh

set -e

VERSION=$1
IMAGE_NAME=$2
IP=$3
PORT=$4

if test -z $VERSION
then
	echo "请输入版本名称!!!"
	exit 1
fi

if test -z $IMAGE_NAME
then
	echo "请输入镜像名称!!!"
	exit 1
fi

if test -z $IP
then
	echo "请输入宿主映射地址!!!"
	exit 1
fi

if test -z $PORT
then
	echo "请输入宿主映射端口!!!"
	exit 1
fi

# 停止并删除正在运行的容器
IMAGE_ID=$(docker ps -a | grep $IMAGE_NAME | head -1 | awk '{print $1}')
if [ $IMAGE_ID != "" ]; then
  docker stop $IMAGE_ID && docker rm $IMAGE_ID
fi

# 运行容器
docker pull $IMAGE_NAME:$VERSION
docker run -d -p $IP:$PORT:$PORT $IMAGE_NAME