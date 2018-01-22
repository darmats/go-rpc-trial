#!/bin/sh

DIR=$(cd $(dirname $0); pwd)
cd ${DIR}

thrift -out ../../ -r --gen go interface/hello.thrift
