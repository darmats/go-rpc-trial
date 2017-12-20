#!/bin/sh

DIR=$(cd $(dirname $0); pwd)
cd ${DIR}

protoc -I proto/ proto/*.proto --go_out=plugins=grpc:./pb
