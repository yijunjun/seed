#!/bin/sh
# protoc --go_out=../pb *.proto
protoc --go_out=plugins=grpc:../pb *.proto