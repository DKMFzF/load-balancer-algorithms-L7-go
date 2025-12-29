package weightedroundrobin

import (
	"log"
	"net/url"
	"sync"
)

type Backend struct {
	URL           *url.URL
	Weight        int
	currentWeight int
}

type WeightedRoundRobin struct {
	Backends    []*Backend
	totalWeight int
	mu          sync.Mutex
}

func New(backends []*Backend) *WeightedRoundRobin {
	total := 0

	for _, b := range backends {
		total += b.Weight
	}

	return &WeightedRoundRobin{
		Backends:    backends,
		totalWeight: total,
	}
}

func (w *WeightedRoundRobin) Next() *url.URL {
	w.mu.Lock()
	defer w.mu.Unlock()

	var selected *Backend

	for _, b := range w.Backends {
		b.currentWeight += b.Weight

		if selected == nil || b.currentWeight > selected.currentWeight {
			selected = b
		}
	}

	log.Printf("-> %s (cw=%d)", selected.URL.Host, selected.currentWeight)

	selected.currentWeight -= w.totalWeight
	return selected.URL
}
