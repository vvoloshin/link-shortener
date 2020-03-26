package handlers

import (
	"github.com/vvoloshin/link-shortener/crypto"
	"github.com/vvoloshin/link-shortener/storage"
	"net/http"
)

//кодировка, сохранение строки, возврат хеша
func EncodeUrl(s storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if url := r.PostFormValue("url"); url != "" {
			hashed := crypto.Hash(url)
			s.Save(hashed, url)
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
		if url := r.PostFormValue("url"); url != "" {
			if val, err := s.Read(url); err == nil {
				w.Write([]byte(val))
			} else {
				w.WriteHeader(400)
				w.Write([]byte("requested url not found"))
			}
		}
	}
	return http.HandlerFunc(handleFunc)
}
