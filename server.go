package main

import (
    "fmt"
    "log"
    "net/http"
    "crypto/boring"
)

var greeting string = "Hello"

func handleRequest(w http.ResponseWriter, r *http.Request) {
    if boring.Enabled() {
        fmt.Fprintf(w, "%s OpenShift! (fips-ready)\n", greeting)
    } else {
        fmt.Fprintf(w, "%s OpenShift!\n", greeting)
   }
}

func main() {
    http.HandleFunc("/", handleRequest)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
