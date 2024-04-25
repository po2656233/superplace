#!/bin/bash

# 更新系统软件包
#sudo yum update -y

# 安装 Go 语言环境（如果尚未安装）
#sudo yum install golang -y

# 创建 NATS 用户和组
sudo groupadd nats
sudo useradd -r -g nats -s /sbin/nologin -c "NATS Server" nats

# 下载 NATS Server 二进制文件
curl -L https://github.com/nats-io/nats-server/releases/download/v2.9.10/nats-server-v2.9.10-linux-amd64.tar.gz -o nats-server.tar.gz

# 解压并安装 NATS Server
sudo tar -xzf nats-server.tar.gz -C /usr/local/bin
sudo chown nats:nats /usr/local/bin/nats-server-v2.9.10-linux-amd64/nats-server
sudo ln -s /usr/local/bin/nats-server-v2.9.10-linux-amd64/nats-server /usr/local/bin/nats-server

# 创建 NATS 配置文件
sudo mkdir /etc/nats
sudo chown nats:nats /etc/nats
sudo cp /usr/local/bin/nats-server-v2.9.10-linux-amd64/nats-server.conf /etc/nats/nats-server.conf
sudo chown nats:nats /etc/nats/nats-server.conf

# 创建 NATS 服务文件
cat << EOF | sudo tee /etc/systemd/system/nats.service
[Unit]
Description=NATS Server
After=network.target

[Service]
User=nats
Group=nats
ExecStart=/usr/local/bin/nats-server -c /etc/nats/nats-server.conf
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
EOF

# 启动 NATS 服务
sudo systemctl daemon-reload
sudo systemctl start nats

# 设置 NATS 服务开机自启
sudo systemctl enable nats

# 清理安装文件
rm nats-server.tar.gz

# 输出 NATS 服务状态
sudo systemctl status nats
