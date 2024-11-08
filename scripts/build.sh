#!/bin/bash

# 发生错误则退出
set -e

SCRIPTS_DIR=$(dirname "$(readlink -f "$0")")
#echo "SCRIPTS_DIR is "$SCRIPTS_DIR

ROOT_DIR=$(dirname "$SCRIPTS_DIR")
#echo "ROOT_DIR is "$ROOT_DIR

# 构建 mac 可执行文件
go build -o $ROOT_DIR/output/mac_main $ROOT_DIR/cmd/main.go

# 构建 linux 可执行文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $ROOT_DIR/output/linux_main $ROOT_DIR/cmd/main.go

cp $ROOT_DIR/scripts/control.sh $ROOT_DIR/output/
cp $ROOT_DIR/config/config.yml.template $ROOT_DIR/output/config.yml

echo "success!!!"