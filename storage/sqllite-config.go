package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type SQLite struct {
	Name   string
	Driver string
}

func init() {
	_, err := os.Stat(".\\sqlite\\base.db")
	if os.IsNotExist(err) {
		log.Println("database empty, create it")
		os.Mkdir(".\\sqlite", 0755)
		os.Create(".\\sqlite\\base.db")
	}
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
