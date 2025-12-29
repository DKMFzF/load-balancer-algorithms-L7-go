package roundrobin

import (
	"net/url"
	"sync/atomic"
)

type RoundRobineBalancer struct {
	backends []*url.URL
	counter  uint64
}

func (rr *RoundRobineBalancer) NextBackend() *url.URL {
	index := atomic.AddUint64(&rr.counter, 1)
	return rr.backends[index%uint64(len(rr.backends))]
}
