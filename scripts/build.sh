#!/bin/bash

# 发生错误则退出
set -e

SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
#echo "SCRIPT_DIR is "$SCRIPT_DIR

PROJECT_DIR=$(dirname "$SCRIPT_DIR")
#echo "PROJECT_DIR is "$PROJECT_DIR

echo "mkdir -p $PROJECT_DIR/output"
mkdir -p $PROJECT_DIR/output

echo "mkdir -p $PROJECT_DIR/output/configs"
mkdir -p $PROJECT_DIR/output/configs

cp $PROJECT_DIR/scripts/control.sh $PROJECT_DIR/output/control.sh
cp $PROJECT_DIR/configs/config.yml.template $PROJECT_DIR/output/configs/config.yml

# 构建 mac 可执行文件
go build -o $PROJECT_DIR/output/mac_main $PROJECT_DIR/cmd/main.go

# 构建 linux 可执行文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $PROJECT_DIR/output/linux_main $PROJECT_DIR/cmd/main.go

echo "success!!!"