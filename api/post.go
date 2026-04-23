package api

import (
	"fmt"
	"net/http"
	"strings"
)

func Post(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path

	if path == "/post" {

		if req.Method == http.MethodGet {

		}

		if req.Method == http.MethodPost {

		}
	}

	if strings.HasPrefix(path, "/post/") {
		id := strings.TrimPrefix(path, "/post/")

		if id == "" {
			fmt.Print("Error Not found id")
		}
	}
}
