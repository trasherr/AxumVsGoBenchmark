package main

import (
	"net/http"
)

func setJsonResponseHeader(w http.ResponseWriter) {

	// Set the header content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
