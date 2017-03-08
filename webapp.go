package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HomeHandler)
    panic(http.ListenAndServe(":8000", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><head><style>body{ background-color: green;}</style></head><body></body></html>")
	w.WriteHeader(http.StatusOK)
}
