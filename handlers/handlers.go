package handlers

import (
	"github.com/vvoloshin/link-shortener/crypto"
	"github.com/vvoloshin/link-shortener/storage"
	"net/http"
)

//кодировка, сохранение строки, возврат хеша
func EncodeUrl(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if !validRequest(w, r, "POST") {
			return
		}
		if rawUrl := r.PostFormValue("url"); rawUrl != "" {
			hashed := crypto.Hash(rawUrl)
			s.Save(hashed, rawUrl)
			w.WriteHeader(201)
			w.Write([]byte(hashed))
		} else {
			w.WriteHeader(204)
			w.Write([]byte("not specified body with `url` parameter"))
		}
	}
	return http.HandlerFunc(handleFunc)
}

//поиск по хешу, возврат оригинальной строки
//rawUrl - оригинальная строка, не закодированная, хранится в Storage в качестве `value`
//hashed - хэш от rawUrl, хранится в Storage в качестве `key`
func DecodeUrl(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if !validRequest(w, r, "POST") {
			return
		}
		if hashed := r.PostFormValue("url"); hashed != "" {
			if rawUrl, err := s.Read(hashed); err == nil {
				w.Write([]byte(rawUrl))
			} else {
				w.WriteHeader(400)
				w.Write([]byte("requested url not found in Storage"))
			}
		} else {
			w.WriteHeader(400)
			w.Write([]byte("key-url not found in request"))
		}
	}
	return http.HandlerFunc(handleFunc)
}

func Redirect(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if !validRequest(w, r, "POST") {
			return
		}
		if hashed := r.PostFormValue("url"); hashed != "" {
			if rawUrl, err := s.Read(hashed); err == nil {
				http.Redirect(w, r, rawUrl, http.StatusSeeOther)
			} else {
				w.WriteHeader(400)
				w.Write([]byte("requested url not found"))
			}
		} else {
			w.WriteHeader(400)
			w.Write([]byte("key-url not found in request"))
		}
	}
	return http.HandlerFunc(handleFunc)
}

func validRequest(w http.ResponseWriter, r *http.Request, m string) bool {
	if r.Method != m {
		w.WriteHeader(400)
		w.Write([]byte("method " + r.Method + " not allowed"))
		return false
	}
	return true
}
