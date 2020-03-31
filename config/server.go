package config

import "github.com/vvoloshin/link-shortener/storage"

type Server struct {
	Port    string
	Storage storage.Storage
}

func NewDefaultServer(port, file, driver string) *Server {
	return &Server{
		Port:    port,
		Storage: storage.NewSQLite(file, driver),
	}
}
