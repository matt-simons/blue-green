package main

import "net/http"

func main() {
    fs := http.FileServer(http.Dir("/go/src/github.com/matt-simons/blue-green/static"))
    http.Handle("/", fs)
    panic(http.ListenAndServe(":8000", nil))
}

