package main

import (
	roundrobin "balancer/internal/roundRobin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	var urls []*url.URL
	for _, b := range backends {
		u, _ := url.Parse(b)
		urls = append(urls, u)
	}

	rr := &roundrobin.Roundrobinbalancer{Backends: urls}

	handler := func(w http.ResponseWriter, r *http.Request) {
		target := rr.Nextbackend()
		proxy := httputil.NewSingleHostReverseProxy(target)

		log.Printf("→ %s\n", target.Host)
		proxy.ServeHTTP(w, r)
	}

	http.HandleFunc("/", handler)

	log.Println("Load balancer listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
