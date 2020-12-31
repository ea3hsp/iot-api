#!/usr/bin/env zsh
echo "*** PROTOBUFFER GENERATING ***"
protoc --proto_path=$GOPATH/src:. --gofast_out=plugins=grpc:. ./pb/*.proto
protoc --proto_path=$GOPATH/src:. --nanopb_out=. ./pb/domo.proto
echo "*** DONE PROTOBUFFER GENERATING ***"