package main

import (
	"fmt"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "params：%s\n", req.URL.RawQuery)
}
