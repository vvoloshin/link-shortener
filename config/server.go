package config

import "github.com/vvoloshin/link-shortener/storage"

type Server struct {
	Port    string
	Storage storage.Storage
}

func NewDefault() *Server {
	return &Server{
		Port:    ":8080",
		Storage: storage.InMemStorage{Store: map[string]string{}},
	}
}
