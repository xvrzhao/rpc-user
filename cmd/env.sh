#!/bin/bash
# 导入环境变量

if [[ ! -f "../.env" ]]; then
  echo "error: no .env file"
  exit
fi

while read line; do
  if [[ $line != \#* && $line != "" ]]; then
    export ${line}
  fi
done < ../.env