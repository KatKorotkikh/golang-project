package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux)
}

func HelloHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Hello World!")
}
