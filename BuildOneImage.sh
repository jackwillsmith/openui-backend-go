#!/bin/bash

# 根据名称匹配要编译的应用路径
APP_NAME="$1"

# prepare env
export WORKSPACE=cloud
export VERSION="$(cat ./version)"
export FinishImage="eilinge/$APP_NAME"
export TZ=UTC-8
export TIME_ID=$(date +%Y-%m-%d-%H-%M-%S)
export IMGAE_NAME="${FinishImage}:$VERSION"

# start build
docker build -t $IMGAE_NAME --build-arg AppDir=$APP_NAME -f ./Dockerfile .
