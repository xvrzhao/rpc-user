#!/bin/bash
# 启动应用

project="rpc-user"

cd $(dirname $0)
source env.sh

go build -o ${project} ../*.go
./${project}