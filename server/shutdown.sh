#!/bin/sh

DIR=$(cd $(dirname $0); pwd)
cd ${DIR}

PS=$(ps -ef | grep ./go-rpc-trial-backend | grep -v "grep" | awk '{print $2;}')
if [ -n "${PS}" ]; then
  kill ${PS}
fi

PS=$(ps -ef | grep ./go-rpc-trial-proxy | grep -v "grep" | awk '{print $2;}')
if [ -n "${PS}" ]; then
  kill ${PS}
fi
