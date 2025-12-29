package sidecar

import (
	"balancer/internal/metrics"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	metrics.Init()
	http.Handle("/metrics", promhttp.Handler())
}
