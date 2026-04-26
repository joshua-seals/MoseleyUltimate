#!/bin/bash
dnf install -y docker curl
systemctl enable docker
systemctl start docker
usermod -aG docker ec2-user

mkdir -p /usr/libexec/docker/cli-plugins
curl -SL https://github.com/docker/compose/releases/latest/download/docker-compose-linux-x86_64 -o /usr/libexec/docker/cli-plugins/docker-compose
chmod +x /usr/libexec/docker/cli-plugins/docker-compose

cd /home/ec2-user
git clone https://github.com/joshua-seals/NextonFrisbeeClub.git
chown -R ec2-user:ec2-user NextonFrisbeeClub 

sudo systemctl daemon-reexec
sudo systemctl daemon-reload
sudo systemctl enable nextonfrisbee-app.service