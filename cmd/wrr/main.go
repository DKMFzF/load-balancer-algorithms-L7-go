package main

import (
	wrr "balancer/internal/weightedRoundRobin"
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
	rr := wrr.New([]*wrr.Backend{
		{URL: mustParse("http://localhost:8081"), Weight: 5},
		{URL: mustParse("http://localhost:8082"), Weight: 3},
		{URL: mustParse("http://localhost:8083"), Weight: 1},
	})

	handler := func(w http.ResponseWriter, r *http.Request) {
		target := rr.Next()
		proxy := httputil.NewSingleHostReverseProxy(target)

		//log.Printf("-> %s\n", target.Host)
		proxy.ServeHTTP(w, r)
	}

	http.HandleFunc("/", handler)

	log.Printf("Weighted RR LB on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
