#!/bin/bash

version=$(go run ./cmd/apiping -version | awk '{ print $2 }' | awk -F= '{ print $2 }')

echo version=$version

docker build --no-cache \
    -t udhos/apiping:latest \
    -t udhos/apiping:$version \
    -f docker/Dockerfile .

echo "push: docker push udhos/apiping:$version; docker push udhos/apiping:latest"
