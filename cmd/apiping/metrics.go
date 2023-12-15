package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	latencySpringServer *prometheus.HistogramVec
	latencySpringClient *prometheus.HistogramVec
}

var (
	dimensionsSpring = []string{"method", "status", "uri"}
)

const (
	latencySpringNameServer = "http_server_requests_seconds"
	latencySpringNameClient = "http_client_requests_seconds"
)

func newMetrics(namespace string, latencyBucketsServer, latencyBucketsClient []float64) *metrics {
	const me = "newMetrics"

	//
	// latency server
	//

	latencySpringServer := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      latencySpringNameServer,
			Help:      "Spring-like server request duration in seconds.",
			Buckets:   latencyBucketsServer,
		},
		dimensionsSpring,
	)

	if err := prometheus.Register(latencySpringServer); err != nil {
		log.Fatalf("%s: server latency was not registered: %s", me, err)
	}

	//
	// latency client
	//

	latencySpringClient := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      latencySpringNameClient,
			Help:      "Spring-like client request duration in seconds.",
			Buckets:   latencyBucketsClient,
		},
		dimensionsSpring,
	)

	if err := prometheus.Register(latencySpringClient); err != nil {
		log.Fatalf("%s: client latency was not registered: %s", me, err)
	}

	//
	// all metrics
	//

	m := &metrics{
		latencySpringServer: latencySpringServer,
		latencySpringClient: latencySpringClient,
	}

	return m
}

func (m *metrics) recordLatencyServer(method, status, path string, latency time.Duration) {
	m.latencySpringServer.WithLabelValues(method, status, path).Observe(float64(latency) / float64(time.Second))
}

func (m *metrics) recordLatencyClient(method, status, path string, latency time.Duration) {
	m.latencySpringClient.WithLabelValues(method, status, path).Observe(float64(latency) / float64(time.Second))
}
