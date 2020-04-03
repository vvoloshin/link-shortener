package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vvoloshin/link-shortener/util"
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
	util.CheckErrorVerb(err, "can't connect to database: "+file)
	err = db.Ping()
	util.CheckErrorVerb(err, "can't ping database: "+file)
	return db
}
