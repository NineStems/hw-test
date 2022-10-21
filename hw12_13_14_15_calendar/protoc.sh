#!/bin/bash
set -e

module="github.com/hw-test/hw12_13_14_15_calendar/api/calendar"
#Find all dirs with .proto files in them
for proto in ./proto/*.proto; do
    #echo "regenerating generated protobuf code for ${proto}"
    #protoc --proto_path .\
    #       --go-grpc_out=./api/calendar --go-grpc_opt=module=${module}\
    #       --go_out=./api/calendar --go_opt=module=${module}\
    #       ${proto}
    echo "creating reverse proxy protobuf code for ${proto}"
    protoc --grpc-gateway_out=. ${proto}
    #protoc -I . --internalgrpc-gateway_out=. \
    # --internalgrpc-gateway_opt logtostderr=true \
    # --internalgrpc-gateway_opt paths=source_relative \
     #${proto}
done