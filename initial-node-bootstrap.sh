#!/bin/bash
 curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
 add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
 curl -L https://github.com/docker/compose/releases/download/1.18.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
 apt-get update
 apt-get install -y docker-ce
 apt-get install -y docker-compose
 docker-compose build
 docker-compose up -d
 sleep 10
./conf/postgres/sql.sh

