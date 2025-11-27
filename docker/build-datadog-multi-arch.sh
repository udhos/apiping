#!/bin/bash

app=apiping

version=$(go run ./cmd/$app -version | awk '{ print $2 }' | awk -F= '{ print $2 }')

dd=-datadog

echo version=$version

docker build --no-cache \
   --push \
   --build-arg GOARCH=amd64 \
   -t udhos/$app:${version}-amd64${dd} \
   -f docker/Dockerfile.datadog .

docker build --no-cache \
   --push \
   --build-arg GOARCH=arm64 \
   -t udhos/$app:${version}-arm64${dd} \
   -f docker/Dockerfile.datadog .

manifest() {
   local tag="$1"
   docker manifest create udhos/$app:$tag${dd} udhos/$app:${version}-amd64${dd} udhos/$app:${version}-arm64${dd}
   docker manifest annotate --arch amd64 udhos/$app:$tag${dd} udhos/$app:${version}-amd64${dd}
   docker manifest annotate --arch arm64 udhos/$app:$tag${dd} udhos/$app:${version}-arm64${dd}
}

manifest $version
manifest latest

echo push:

cat >docker-push-datadog.sh << EOF
docker manifest push udhos/$app:$version${dd}
docker manifest push udhos/$app:latest${dd}
EOF

chmod a+rx docker-push-datadog.sh
echo docker-push-datadog.sh:
cat docker-push-datadog.sh
