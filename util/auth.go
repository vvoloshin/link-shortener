package util

import (
	"net/http"
)

func IsAuthenticated(c *Config, w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get(c.ApiSecure.ApiHeader) == c.ApiSecure.ApiKey {
		return true
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("not specified authentication"))
	return false
}
