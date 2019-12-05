#!/bin/bash

set -e

NAMESPACE=fattigerlee
IMAGE_NAME=test_cicd
REGISTRY_URL=registry.cn-hangzhou.aliyuncs.com

echo "构建镜像..."

#版本号
VERSION=$1

if test -z $VERSION
then
	echo "请输入版本号"
	exit 1
fi

echo "${VERSION}" > ../../configs/${IMAGE_NAME}/VERSION

#生成dockfile文件
cat>Dockerfile<<EOF
FROM alpine

WORKDIR /root/

COPY ${IMAGE_NAME} ./app/

CMD ["./app/${IMAGE_NAME}"]
EOF

#构建镜像
docker build -t ${IMAGE_NAME} .

echo "构建镜像完成..."

echo "推送镜像..."

docker login --username="505179140@qq.com" --password="C317t#VrjAb*r%uv" ${REGISTRY_URL}

docker tag $(docker images | grep ${IMAGE_NAME} | head -1 | awk '{print $3}') ${REGISTRY_URL}/${NAMESPACE}/${IMAGE_NAME}:${VERSION}
docker push ${REGISTRY_URL}/${NAMESPACE}/${IMAGE_NAME}:${VERSION}

docker tag $(docker images | grep ${IMAGE_NAME} | head -1 | awk '{print $3}') ${REGISTRY_URL}/${NAMESPACE}/${IMAGE_NAME}:latest
docker push ${REGISTRY_URL}/${NAMESPACE}/${IMAGE_NAME}:latest

echo "推送镜像完成..."

#删除不要的文件
rm -rf ${IMAGE_NAME}
rm -rf Dockerfile
docker image prune -a -f