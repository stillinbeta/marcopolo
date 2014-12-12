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

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Fatal(http.ListenAndServe(":" + port, nil))
}
