package main

import (
    "fmt"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("images"))
    http.HandleFunc("/", homeHandler)
    panic(http.ListenAndServe(":8000", nil))
}

