# apiping

apiping

# Build

```bash
git clone https://github.com/udhos/apiping
cd apiping
./build.sh
```

# Run and test

```
apiping
```

```bash
curl localhost:8080/ping
ok
```

# Docker

Docker hub:

https://hub.docker.com/r/udhos/apiping

Run from docker hub:

```
docker run -p 8080:8080 --rm udhos/apiping:0.0.0
```

Build recipe:

```
./docker/build.sh

docker push udhos/apiping:0.0.0
```
