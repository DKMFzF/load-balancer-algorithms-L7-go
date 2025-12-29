package main

import (
	lc "balancer/internal/leastconn"
	"balancer/internal/metrics"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func mustParse(raw string) *url.URL {
	u, err := url.Parse(raw)
	if err != nil {
		panic(err)
	}
	return u
}

func main() {
	metrics.Init()

	lb := lc.New([]*lc.Backend{
		{URL: mustParse("http://localhost:8081/")},
		{URL: mustParse("http://localhost:8082/")},
		{URL: mustParse("http://localhost:8083/")},
	})

	handler := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		backend := lb.Acquire()

		cw := &lc.ContingResponseWriter{
			ResponseWriter: w,
			OnClose: func() {
				lb.Release(backend)
			},
		}

		proxy := httputil.NewSingleHostReverseProxy(backend.URL)

		log.Printf("-> %s (conn=%d)", backend.URL.Host, backend.Connections)

		proxy.ServeHTTP(cw, r)
		cw.Finish()

		metrics.RequestsTotal.
			WithLabelValues(backend.URL.Host).
			Inc()

		metrics.RequestDuration.
			WithLabelValues(backend.URL.Host).
			Observe(time.Since(start).Seconds())
	}

	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Least Connections on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
