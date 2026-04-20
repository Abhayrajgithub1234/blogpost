package api

import (
	"net/http"
)

func Post(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
}
