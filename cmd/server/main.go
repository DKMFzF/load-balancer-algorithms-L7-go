package main

import (
	"fmt"
	"log"
	"strconv"

	//"math/rand"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	name := os.Getenv("NAME")
	delay := os.Getenv("DELAY")

	if port == "" || name == "" || delay == "" {
		log.Fatal("Port or Name not found in env")
	}

	var counter int64
	intDelay, err := strconv.Atoi(delay)
	if err != nil {
		log.Fatalf("%v", err)
	}
	//minDelay := 10
	//maxDelay := 200

	handler := func(w http.ResponseWriter, r *http.Request) {
		//randomDelay := rand.Intn(maxDelay-minDelay+1) + minDelay
		n := atomic.AddInt64(&counter, 1)
		time.Sleep(time.Duration(intDelay) * time.Millisecond)
		fmt.Fprintf(w, "pong from %s | count=%d | delay=%d\n", name, n, delay)
	}

	http.HandleFunc("/", handler)

	log.Printf("%s listening on %s\n", name, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
