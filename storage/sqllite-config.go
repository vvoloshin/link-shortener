package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SQLite struct {
	Name   string
	Driver string
	*sql.DB
}

func NewSQLite(file, driver string) *SQLite {
	return &SQLite{
		Name:   file,
		Driver: driver,
		DB:     connect(file, driver),
	}
}

func connect(file, driver string) *sql.DB {
	db, err := sql.Open(driver, file)
	if err != nil {
		log.Fatal("can't connect to database: ", file)
	}
	err = db.Ping()
	if err != nil {
		log.Println("can't ping to database: ", file)
		log.Fatal(err)
	}
	return db
}
