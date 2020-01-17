#!/bin/bash
# 运行所有单元测试

project="rpc-user"

cd $(dirname $0)
source env.sh
export RUN_MODE=test

cd ..
go test -v -race -count=1 ./...