#!/bin/bash
sudo yum update -y
sudo amazon-linux-extras install docker
sudo service docker start
sudo systemctl enable docker

sudo docker run --rm -d --privileged -p 500:500/udp -p 4500:4500/udp -p 1701:1701/tcp \
  -e PSK='yourkey' \
  -e USERS='user:password' \
  siomiz/softethervpn
