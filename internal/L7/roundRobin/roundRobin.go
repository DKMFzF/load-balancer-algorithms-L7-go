package roundrobin

import (
	"net/url"
	"sync/atomic"
)

type Roundrobinbalancer struct {
	Backends []*url.URL
	counter  uint64
}

func (rr *Roundrobinbalancer) Nextbackend() *url.URL {
	index := atomic.AddUint64(&rr.counter, 1)
	return rr.Backends[index%uint64(len(rr.Backends))]
}
