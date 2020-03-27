package handlers

import (
	"github.com/vvoloshin/link-shortener/crypto"
	"github.com/vvoloshin/link-shortener/storage"
	"net/http"
)

//кодировка, сохранение строки, возврат хеша
func EncodeUrl(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
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
func DecodeUrl(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if hashed := r.PostFormValue("url"); hashed != "" {
			if val, err := s.Read(hashed); err == nil {
				w.Write([]byte(val))
			} else {
				w.WriteHeader(400)
				w.Write([]byte("requested url not found"))
			}
		}
	}
	return http.HandlerFunc(handleFunc)
}

func Redirect(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if hashed := r.PostFormValue("url"); hashed != "" {
			if url, err := s.Read(hashed); err == nil {
				http.Redirect(w, r, url, http.StatusSeeOther)
			} else {
				w.WriteHeader(400)
				w.Write([]byte("requested url not found"))
			}
		}
	}
	return http.HandlerFunc(handleFunc)
}
