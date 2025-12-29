package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	name := os.Getenv("NAME")

	if port == "" || name == "" {
		log.Fatal("Port or Name not found in env")
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong from %s\n", name)
	}

	http.HandleFunc("/", handler)

	log.Printf("%s listening on %s\n", name, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
