#!/bin/sh
git pull origin
sudo docker build -t osp-base .
sudo docker stop osp-base
sudo docker rm osp-base
sudo docker run -d --name osp-base --restart=always -p 8088:8088 -v /Storage/allure-server/projects:/mnt --network osp-network --log-opt max-size=30m --log-opt max-file=3 osp-base:latest
