package leastconn

import (
	"balancer/internal/metrics"
	"net/url"
	"sync"
)

type Backend struct {
	URL         *url.URL
	Connections int64
}

type LeastConnections struct {
	Backends []*Backend
	mu       sync.Mutex
}

func New(backends []*Backend) *LeastConnections {
	return &LeastConnections{
		Backends: backends,
	}
}

func (lc *LeastConnections) Acquire() *Backend {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	var selected *Backend

	for _, b := range lc.Backends {
		if selected == nil || b.Connections < selected.Connections {
			selected = b
		}
	}

	selected.Connections++

	metrics.ActiveConnections.WithLabelValues(selected.URL.Host).Inc()

	return selected
}

func (lc *LeastConnections) Release(b *Backend) {
	lc.mu.Lock()
	b.Connections--
	lc.mu.Unlock()

	metrics.ActiveConnections.WithLabelValues(b.URL.Host).Dec()
}
