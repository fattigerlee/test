#!/bin/bash

echo "这是一个测试脚本"

version=$1

if test -z $version
then
	echo "请输入版本号"
	exit 1
fi

app=$(docker images | grep nginx | head -n 1 | awk '{print $1}')

if test -z $app
then
	echo "app名称生成失败"
	exit 1
fi

echo "程序信息:${app}:${version}"
echo '测试结束'