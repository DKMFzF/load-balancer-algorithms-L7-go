package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	ActiveConnections = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "lb",
			Name:      "active_connections",
			Help:      "Active connections per backend",
		},
		[]string{"backend"},
	)

	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "lb",
			Name:      "requests_total",
			Help:      "Total requests",
		},
		[]string{"backend"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "lb",
			Name:      "request_duration_seconds",
			Help:      "Request latency",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"backend"},
	)
)

func Init() {
	prometheus.MustRegister(
		ActiveConnections,
		RequestsTotal,
		RequestDuration,
	)
}
