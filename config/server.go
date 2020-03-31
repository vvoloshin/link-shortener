package config

import "github.com/vvoloshin/link-shortener/storage"

type Server struct {
	Port    string
	Storage storage.Storage
}

func NewDefaultSQLite() *Server {
	//todo: добавить инициализацию файла для sqlite
	return &Server{
		Port: ":8080",
		Storage: storage.SQLite{
			Name:   ".\\sqlite\\base.db",
			Driver: "sqlite3",
		},
	}
}
