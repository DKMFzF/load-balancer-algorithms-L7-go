package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	name := os.Getenv("NAME")

	if port == "" || name == "" {
		log.Fatal("Port or Name not found in env")
	}

	var counter int64
	minDelay := 10
	maxDelay := 1000

	handler := func(w http.ResponseWriter, r *http.Request) {
		randomDelay := rand.Intn(maxDelay-minDelay+1) + minDelay
		n := atomic.AddInt64(&counter, 1)
		time.Sleep(time.Duration(randomDelay) * time.Millisecond)
		fmt.Fprintf(w, "pong from %s | count=%d | delay=%d\n", name, n, randomDelay)
	}

	http.HandleFunc("/", handler)

	log.Printf("%s listening on %s\n", name, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
