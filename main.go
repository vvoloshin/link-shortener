package main

import (
	"github.com/vvoloshin/link-shortener/config"
	"github.com/vvoloshin/link-shortener/handlers"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	driver = "sqlite3"
	port   = ":8080"
)

var (
	file = filepath.FromSlash("sqlite\\base.db")
	dir  = filepath.FromSlash("sqlite")
)

func init() {
	if isFileNotExist(file) {
		log.Println("database empty, created it")
		createFile(dir, file)
	}
	log.Println("found existing database file")
}

func main() {
	server := config.NewDefaultServer(port, file, driver)
	err := server.Storage.InitTable()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/encode", handlers.EncodeUrl(server.Storage))
	http.Handle("/decode", handlers.DecodeUrl(server.Storage))
	http.Handle("/redirect", handlers.Redirect(server.Storage))
	log.Println("starts server at port: " + server.Port)
	err = http.ListenAndServe(server.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createFile(d, f string) {
	os.Mkdir(d, 0755)
	os.Create(f)
}

func isFileNotExist(f string) bool {
	_, err := os.Stat(f)
	return os.IsNotExist(err)
}
