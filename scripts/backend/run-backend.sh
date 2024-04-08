#!/bin/bash

BASE_DIR=$(dirname "$0")/../..
GO_IMAGE_NAME=ustar-backend

docker stop ${GO_IMAGE_NAME}
docker container rm ${GO_IMAGE_NAME}

docker build . -t ${GO_IMAGE_NAME} -f "${BASE_DIR}"/Dockerfile

docker run -d --name ${GO_IMAGE_NAME} \
  -p 8080:8080 \
  ${GO_IMAGE_NAME}
