#!/bin/bash
 set -e 
 echo "Adding docker repo and performing installation"
 curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
 add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
 apt-get update
 echo "Installing Docker and docker-compose and psql client"
 apt-get install -y docker-ce
 apt-get install -y docker-compose
 apt-get install postgresql-client
 echo "Building docker images"
 docker-compose build
 echo "starting docker instances"
 docker-compose up -d
 echo "sleeping for 10 Sec"
 sleep 10
 echo "creating DB Tables. see docker-compose file for the password "
./conf/postgres/sql.sh

