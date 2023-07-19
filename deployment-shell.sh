#!/bin/bash

#停止服务
docker stop Mercury-web-api
docker stop Mercury-web-rpc
docker stop Mercury-job

#删除容器
docker rm Mercury-web-api
docker rm Mercury-web-rpc
docker rm Mercury-job

#删除镜像
docker rmi Mercury-web-api:v1
docker rmi Mercury-web-rpc:v1
docker rmi Mercury-job:v1

#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t austin-job:v1 -f app/Mercury-job/Dockerfile .
docker build -t austin-web-rpc:v1 -f app/Mercury-web/rpc/Dockerfile .
docker build -t austin-web-api:v1 -f app/Mercury-web/api/Dockerfile .

#启动服务
docker run -itd --net=host --name=Mercury-job Mercury-job:v1
docker run -itd --net=host --name=Mercury-web-rpc Mercury-web-rpc:v1
docker run -itd --net=host --name=Mercury-web-api Mercury-web-api:v1

