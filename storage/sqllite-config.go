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

const (
	file = ".\\sqlite\\base.db"
	dir  = ".\\sqlite"
)

func init() {
	if isFileNotExist(file) {
		log.Println("database empty, create it")
		createFile(dir, file)
	}
	log.Println("found existing database file")
}

func (s SQLite) connect() *sql.DB {
	db, err := sql.Open(s.Driver, s.Name)
	if err != nil {
		log.Fatal("can't connect to database: ", s.Name)
	}
	err = db.Ping()
	if err != nil {
		log.Println("can't ping to database: ", s.Name)
		log.Fatal(err)
	}
	return db
}

func createFile(d, f string) {
	os.Mkdir(d, 0755)
	os.Create(f)
}

func isFileNotExist(f string) bool {
	_, err := os.Stat(f)
	return os.IsNotExist(err)
}
