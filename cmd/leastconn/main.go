package main

import (
	lc "balancer/internal/leastconn"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func mustParse(raw string) *url.URL {
	u, err := url.Parse(raw)
	if err != nil {
		panic(err)
	}
	return u
}

func main() {
	lb := lc.New([]*lc.Backend{
		{URL: mustParse("http://localhost:8081/")},
		{URL: mustParse("http://localhost:8082/")},
		{URL: mustParse("http://localhost:8083/")},
	})

	handler := func(w http.ResponseWriter, r *http.Request) {
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
	}

	http.HandleFunc("/", handler)

	log.Println("Least Connections on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
