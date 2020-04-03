package handlers

import (
	"github.com/vvoloshin/link-shortener/crypto"
	"github.com/vvoloshin/link-shortener/storage"
	"github.com/vvoloshin/link-shortener/util"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"strings"
)

func Processing(c *util.Config, s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if !validRequestMethod(w, r, http.MethodPost) {
			return
		}
		if !util.IsAuthenticated(c, w, r) {
			return
		}
		if !hasContentType(r, "text/plain") {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("not specified or unsupported media-type"))
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		nonEmptyRawUrls := splitToLines(body)
		var hashedUrls []string
		for _, v := range nonEmptyRawUrls {
			hashed := generateVerifiedHash(s)
			s.Save(hashed, v)
			hashedUrls = append(hashedUrls, c.ServerHost.Host+hashed)
		}
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, strings.Join(hashedUrls, "\n"))
	}
	return http.HandlerFunc(handleFunc)
}

func Redirect(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if !validRequestMethod(w, r, http.MethodGet) {
			return
		}
		if hashed := strings.TrimPrefix(r.URL.Path, "/"); hashed != "" {
			if rawUrl, err := s.Read(hashed); err == nil {
				http.Redirect(w, r, rawUrl, http.StatusMovedPermanently)
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("requested url not found"))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	return http.HandlerFunc(handleFunc)
}

func generateVerifiedHash(s storage.Storage) string {
	var hashed string
	for {
		hashed = crypto.GenerateBase58Str()
		stored, _ := s.Read(hashed)
		if stored == "" {
			break
		}
	}
	return hashed
}

func splitToLines(body []byte) []string {
	rawUrlsBody := strings.Split(string(body), "\n")
	var nonEmptyRawUrls []string
	for _, v := range rawUrlsBody {
		if v != "" {
			nonEmptyRawUrls = append(nonEmptyRawUrls, v)
		}
	}
	return nonEmptyRawUrls
}

func validRequestMethod(w http.ResponseWriter, r *http.Request, m string) bool {
	if r.Method != m {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method " + r.Method + " not allowed"))
		return false
	}
	return true
}

func hasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		return false
	}
	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}
