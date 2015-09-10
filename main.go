package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("%v\n", os.Args)
	http.HandleFunc("/marco", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/plain")
		w.Header().Set("X-Marco-Version", "0.0.3")
		fmt.Fprintf(w, "POLO!\n")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/plain")
		fmt.Fprintf(w, "yes good\n")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
