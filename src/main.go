package main

import (
	"net/http"
)

func main() {
	PORT := ":8080"

	http.ListenAndServe(PORT, nil)
}
