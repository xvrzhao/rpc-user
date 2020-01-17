#!/bin/bash
# 启动应用

project="rpc-user"

cd $(dirname $0)
source env.sh

go build -o ${project} ../*.go
if [ $? -ne 0 ]; then
    echo -e "\nstartup.sh: failed to build"
    exit
fi

./${project}