package main

import (
    "net/http"
    "fmt"
    "os"
    "log"
)


func main() {
    http.HandleFunc("/marco", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-type", "text/plain")
        fmt.Fprintf(w, "POLO!\n")
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-type", "text/plain")
        fmt.Fprintf(w, "yes good\n")
    })
    port := os.Getenv("PORT")
    if port == "" {
        port = "80"
    }

    log.Fatal(http.ListenAndServe("0.0.0.0:" + port, nil))
}
