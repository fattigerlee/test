#!/bin/sh

set -e

VERSION=$1
REGISTRY=$2
REGISTRY_NAMESPACE=$3
APP=$4

if test -z $VERSION
then
	echo "请输入版本名称!!!"
	exit 1
fi

if test -z $REGISTRY
then
	echo "请输入镜像仓库!!!"
	exit 1
fi

if test -z $REGISTRY
then
	echo "请输入镜像命名空间!!!"
	exit 1
fi

if test -z $APP
then
	echo "请输入镜像名称!!!"
	exit 1
fi

echo "构建镜像..."

IMAGE_NAME=$REGISTRY/$REGISTRY_NAMESPACE/$APP
docker build -t build_cache -f docker/build/Dockerfile .
docker build -t $IMAGE_NAME -f docker/$APP/Dockerfile .

echo "构建镜像完成..."

echo "推送镜像..."

docker tag $(docker images | grep $IMAGE_NAME | head -1 | awk '{print $3}') $IMAGE_NAME:$VERSION
docker push $IMAGE_NAME:$VERSION

echo "推送镜像完成..."