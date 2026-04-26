#!/bin/bash
ADMIN_PASSWORD="change_this_password"

dnf install -y docker curl git
systemctl enable docker
systemctl start docker
usermod -aG docker ec2-user

mkdir -p /usr/libexec/docker/cli-plugins
curl -SL https://github.com/docker/compose/releases/latest/download/docker-compose-linux-x86_64 -o /usr/libexec/docker/cli-plugins/docker-compose
chmod +x /usr/libexec/docker/cli-plugins/docker-compose

cd /home/ec2-user
git clone https://github.com/joshua-seals/MoseleyUltimate.git
chown -R ec2-user:ec2-user MoseleyUltimate

cd MoseleyUltimate

echo "ADMIN_PASSWORD=${ADMIN_PASSWORD}" > .env
chmod 600 .env
chown ec2-user:ec2-user .env

docker compose up -d --build
