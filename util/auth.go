package util

import (
	"github.com/vvoloshin/link-shortener/config"
	"net/http"
)

func IsAuthenticated(c *config.Config, w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get(c.ApiHeader) == c.ApiKey {
		return true
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("not specified authentication"))
	return false
}
