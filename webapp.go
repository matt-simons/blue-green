package main

import "net/http"

func main() {
    fs := http.FileServer(http.Dir("images"))
    http.Handle("/", fs)
    panic(http.ListenAndServe(":8000", nil))
}

