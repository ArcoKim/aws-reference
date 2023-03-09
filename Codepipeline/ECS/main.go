package main

import (
    "fmt"
    "net/http"
)

func health(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "OK")
}

func color(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "skyblue")
}

func main() {
    http.HandleFunc("/health", health)
    http.HandleFunc("/color", color)
    http.ListenAndServe(":80", nil)
}