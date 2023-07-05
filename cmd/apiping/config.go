package main

import "time"

type config struct {
	addr                        string
	route                       string
	targets                     string
	jaegerURL                   string
	interval                    time.Duration
	metricsAddr                 string
	metricsPath                 string
	metricsNamespace            string
	metricsLatencyBucketsServer []float64
	metricsLatencyBucketsClient []float64
	healthAddr                  string
	healthPath                  string
}

func getConfig() config {
	return config{
		addr:                        envString("ADDR", ":8080"),
		route:                       envString("ROUTE", "/ping"),
		targets:                     envString("TARGETS", `["http://localhost:8080/ping"]`),
		jaegerURL:                   envString("JAEGER_URL", "http://jaeger-collector:14268/api/traces"),
		interval:                    envDuration("INTERVAL", 10*time.Second),
		metricsAddr:                 envString("METRICS_ADDR", ":3000"),
		metricsPath:                 envString("METRICS_PATH", "/metrics"),
		metricsNamespace:            envString("METRICS_NAMESPACE", ""),
		metricsLatencyBucketsServer: envFloat64Slice("METRICS_BUCKETS_LATENCY_SERVER", []float64{0.000005, 0.00001, 0.000025, 0.00005, 0.0001, 0.00025, 0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1}),
		metricsLatencyBucketsClient: envFloat64Slice("METRICS_BUCKETS_LATENCY_CLIENT", []float64{0.0001, 0.00025, 0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, .5, 1}),
		healthAddr:                  envString("HEALTH_ADDR", ":8888"),
		healthPath:                  envString("HEALTH_PATH", "/health"),
	}
}
