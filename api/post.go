package api

import (
	"net/http"
)

func post(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
}
