package util

import "net/http"

const (
	apikey    = "777"
	apiheader = "x-api-key"
)

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get(apiheader) == apikey {
		w.WriteHeader(http.StatusUnauthorized)
		return true
	}
	w.Write([]byte("not specified authentication"))
	return false
}
