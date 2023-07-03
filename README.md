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

# Configuration env vars

```
export ADDR=:8080
export ROUTE=/ping
export TARGETS='["http://localhost:8080/ping"]'
export JAEGER_URL=http://jaeger-collector:14268/api/traces
export INTERVAL=10s
export METRICS_ADDR=:3000
export METRICS_PATH=/metrics
export METRICS_NAMESPACE=""
export METRICS_BUCKETS_LATENCY="0.000005, 0.00001, 0.000025, 0.00005, 0.0001, 0.00025, 0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1"
export METRICS_BUCKETS_LATENCY="0.0001, 0.00025, 0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, .5, 1"
export HEALTH_ADDR=:8888
export HEALTH_PATH=/health
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
