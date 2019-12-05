#!/bin/bash

echo "生成Dockfile"

cat>Dockfile<<EOF
FROM centos

WORKDIR /root/

COPY test /app/
COPY conf /app/conf

EOF
