#!/bin/bash

app=apiping

version=$(go run ./cmd/$app -version | awk '{ print $2 }' | awk -F= '{ print $2 }')

echo version=$version

docker build --no-cache \
   --push \
   --build-arg GOARCH=amd64 \
   -t udhos/$app:${version}-amd64 \
   -f docker/Dockerfile .

docker build --no-cache \
   --push \
   --build-arg GOARCH=arm64 \
   -t udhos/$app:${version}-arm64 \
   -f docker/Dockerfile .

manifest() {
   local tag="$1"
   docker manifest create udhos/$app:$tag udhos/$app:${version}-amd64 udhos/$app:${version}-arm64
   docker manifest annotate --arch amd64 udhos/$app:$tag udhos/$app:${version}-amd64
   docker manifest annotate --arch arm64 udhos/$app:$tag udhos/$app:${version}-arm64
}

manifest $version
manifest latest

echo push:

cat >docker-push.sh << EOF
docker manifest push udhos/$app:$version
docker manifest push udhos/$app:latest
EOF

chmod a+rx docker-push.sh
echo docker-push.sh:
cat docker-push.sh
