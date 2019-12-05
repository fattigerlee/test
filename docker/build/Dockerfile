FROM golang:1.12.13 as build

ENV GOPROXY https://goproxy.cn

WORKDIR /project/cache

COPY go.mod .
COPY go.sum .

RUN go mod download
