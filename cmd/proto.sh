#!/bin/bash
# compile proto file

cd $(dirname $0)/../
protoc -I=. --go_out=plugins=grpc,paths=source_relative:. ./proto/*.proto
