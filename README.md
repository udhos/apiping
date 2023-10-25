[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/udhos/apiping/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/udhos/apiping)](https://goreportcard.com/report/github.com/udhos/apiping)
[![Go Reference](https://pkg.go.dev/badge/github.com/udhos/apiping.svg)](https://pkg.go.dev/github.com/udhos/apiping)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/apiping)](https://artifacthub.io/packages/search?repo=apiping)
[![Docker Pulls apiping](https://img.shields.io/docker/pulls/udhos/apiping)](https://hub.docker.com/r/udhos/apiping)

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
export INTERVAL=20s
export TIMEOUT=15s
export METRICS_ADDR=:3000
export METRICS_PATH=/metrics
export METRICS_NAMESPACE=""
export METRICS_BUCKETS_LATENCY_SERVER="0.000005, 0.00001, 0.000025, 0.00005, 0.0001, 0.00025, 0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1"
export METRICS_BUCKETS_LATENCY_CLIENT="0.0001, 0.00025, 0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, .5, 1"
export HEALTH_ADDR=:8888
export HEALTH_PATH=/health
export OTEL_TRACES_SAMPLER=parentbased_traceidratio
export OTEL_TRACES_SAMPLER_ARG="0.01"
# pick one of OTEL_SERVICE_NAME or OTEL_RESOURCE_ATTRIBUTES
#export OTEL_SERVICE_NAME=mynamespace.apiping
#export OTEL_RESOURCE_ATTRIBUTES='service.name=mynamespace.apiping,key2=value2'

export OTELCONFIG_EXPORTER=jaeger
export OTEL_TRACES_EXPORTER=jaeger
export OTEL_PROPAGATORS=b3multi
export OTEL_EXPORTER_OTLP_ENDPOINT=http://jaeger-collector:14268

export OTELCONFIG_EXPORTER=grpc
export OTEL_TRACES_EXPORTER=otlp
export OTEL_PROPAGATORS=b3multi
export OTEL_EXPORTER_OTLP_ENDPOINT=http://jaeger-collector:4317

export OTELCONFIG_EXPORTER=http
export OTEL_TRACES_EXPORTER=otlp
export OTEL_PROPAGATORS=b3multi
export OTEL_EXPORTER_OTLP_ENDPOINT=http://jaeger-collector:4318
```

# Open Telemetry Exporter Configuration:

https://opentelemetry.io/docs/concepts/sdk-configuration/otlp-exporter-configuration/

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

# Helm chart

You can use the provided helm charts to install apiping in your Kubernetes cluster.

See: https://udhos.github.io/apiping/

## Lint

    helm lint ./charts/apiping --values charts/apiping/values.yaml

## Debug

    helm template ./charts/apiping --values charts/apiping/values.yaml --debug

## Render at server

    helm install my-apiping ./charts/apiping --values charts/apiping/values.yaml --dry-run

## Install

    helm install my-apiping ./charts/apiping --values charts/apiping/values.yaml

    helm list -A
