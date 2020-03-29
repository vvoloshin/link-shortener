package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SQLite struct {
	Name   string
	Driver string
}

func (s SQLite) connect() *sql.DB {
	db, err := sql.Open(s.Driver, s.Name)
	if err != nil {
		log.Fatal("can't connect to database: ", s.Name)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
