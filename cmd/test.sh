#!/bin/bash
# 运行单元测试，等同于 go test 命令

project="rpc-user"

cd $(dirname $0)
source env.sh
export RUN_MODE=test

cd ..
go test $@