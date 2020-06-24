#!/usr/bin/env bash
NAME="inventory_service.proto"
#FOLDERNAME="python"
#mkdir -p $FOLDERNAME
protoc $NAME --go_out=plugins=grpc:.
#python -m grpc_tools.protoc -I. --python_out=./$FOLDERNAME/. --grpc_python_out=./$FOLDERNAME/. $NAME

#for Google's implementation
#protoc -I=. $NAME --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.

# rum `npm install ts-protoc-gen` for the first time in the proto file folder
#for Improbable's implementation
protoc --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:. --ts_out=service=grpc-web:. ./$NAME
