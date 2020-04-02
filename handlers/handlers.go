package handlers

import (
	"github.com/vvoloshin/link-shortener/crypto"
	"github.com/vvoloshin/link-shortener/storage"
	"mime"
	"net/http"
	"strings"
)

const (
	apikey    = "777"
	apiheader = "x-api-key"
)

//кодировка, сохранение строки, возврат хеша
func EncodeUrl(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if !validRequestMethod(w, r, http.MethodPost) {
			return
		}
		if r.Header.Get(apiheader) != apikey {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not specified authentication"))
			return
		}
		if rawUrl := r.PostFormValue("url"); rawUrl != "" {
			var hashed string
			for {
				//ищем уже сохраненные ключи, если находим (коллизия), генерируем заново, иначе - сохраняем
				hashed = crypto.Encode(rawUrl)
				stored, _ := s.Read(hashed)
				if stored == "" {
					break
				}
			}
			s.Save(hashed, rawUrl)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(hashed))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("not specified body with `url` parameter"))
		}
	}
	return http.HandlerFunc(handleFunc)
}

//кодировка, сохранение пакета строк, возврат хешей
func BundleUrl(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if !validRequestMethod(w, r, http.MethodPost) {
			return
		}
		if r.Header.Get(apiheader) != apikey {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not specified or incorrect authentication"))
			return
		}
		if !hasContentType(r, "text/plain") {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("not specified or unsupported media-type"))
			return
		}
		panic("not implemented")
	}
	return http.HandlerFunc(handleFunc)
}

//поиск по хешу, возврат оригинальной строки
//rawUrl - оригинальная строка, не закодированная, хранится в Storage в качестве `value`
//hashed - хэш от rawUrl, хранится в Storage в качестве `key`
func DecodeUrl(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if !validRequestMethod(w, r, http.MethodPost) {
			return
		}
		if hashed := r.PostFormValue("url"); hashed != "" {
			if rawUrl, err := s.Read(hashed); err == nil {
				w.Write([]byte(rawUrl))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("requested url not found in Storage"))
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("key-url not found in request"))
		}
	}
	return http.HandlerFunc(handleFunc)
}

func Redirect(prefix string, s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if !validRequestMethod(w, r, http.MethodGet) {
			return
		}
		if hashed := strings.TrimPrefix(r.URL.Path, prefix); hashed != "" {
			if rawUrl, err := s.Read(hashed); err == nil {
				http.Redirect(w, r, rawUrl, http.StatusMovedPermanently)
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("requested url not found"))
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("key-url not found in request"))
		}
	}
	return http.HandlerFunc(handleFunc)
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
