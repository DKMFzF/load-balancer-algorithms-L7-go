package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

func main() {
	port := os.Getenv("PORT")
	name := os.Getenv("NAME")

	if port == "" || name == "" {
		log.Fatal("Port or Name not found in env")
	}

	var counter int64

	handler := func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt64(&counter, 1)
		fmt.Fprintf(w, "pong from %s | count=%d\n", name, n)
	}

	http.HandleFunc("/", handler)

	log.Printf("%s listening on %s\n", name, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
