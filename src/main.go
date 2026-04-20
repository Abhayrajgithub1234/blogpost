package main

import (
	"net/http"
)

func main() {
	PORT := ":8080"
	http.HandleFunc("/post", post)
	http.ListenAndServe(PORT, nil)
}
