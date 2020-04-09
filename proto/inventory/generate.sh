#!/usr/bin/env bash
NAME="inventory_service.proto"
#FOLDERNAME="python"
#mkdir -p $FOLDERNAME
protoc $NAME --go_out=plugins=grpc:.
#python -m grpc_tools.protoc -I. --python_out=./$FOLDERNAME/. --grpc_python_out=./$FOLDERNAME/. $NAME
