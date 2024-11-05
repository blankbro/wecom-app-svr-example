#!/bin/bash

scripts_dir=$(dirname "$0")
root_dir=$(dirname "$scripts_dir")

# 构建 mac 可执行文件
go build -o $root_dir/output/mac/main $root_dir/cmd/main.go

# 构建 linux 可执行文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $root_dir/output/linux/main $root_dir/cmd/main.go

echo "success!!!"