package main

import (
    "fmt"
    "log"
    "net/http"
)

var greeting string = "Hello"

func handleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s OpenShift!\n", greeting)
}

func main() {
    http.HandleFunc("/", handleRequest)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
