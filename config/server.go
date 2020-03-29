package config

import "github.com/vvoloshin/link-shortener/storage"

type Server struct {
	Port    string
	Storage storage.Storage
}

func NewDefaultInMem() *Server {
	return &Server{
		Port: ":8080",
		Storage: storage.InMemStorage{
			Store: map[string]string{},
		},
	}
}

func NewDefaultSQLite() *Server {
	//todo: добавить инициализацию файла для sqlite
	return &Server{
		Port: ":8080",
		Storage: storage.SQLite{
			Name:   "c:\\sqlite\\file.db",
			Driver: "sqlite3",
		},
	}
}
