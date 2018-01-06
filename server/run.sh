#!/bin/sh

DIR=$(cd $(dirname $0); pwd)
cd ${DIR}

# backend
cd backend/
go build -i -o ../go-rpc-trial-backend
cd ../

# proxy
cd proxy/
go build -i -o ../go-rpc-trial-proxy
cd ../


./shutdown.sh
nohup ./go-rpc-trial-backend >> go-rpc-trial-backend.log 2>&1 < /dev/null &
nohup ./go-rpc-trial-proxy >> go-rpc-trial-proxy.log 2>&1 < /dev/null &
