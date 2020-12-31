#!/usr/bin/env zsh
echo "*** PROTOBUFFER GENERATING ***"
protoc --proto_path=$GOPATH/src:. --gofast_out=plugins=grpc:. ./pb/*.proto
echo "*** DONE PROTOBUFFER GENERATING ***"