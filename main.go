package main

import (
	"github.com/vvoloshin/link-shortener/handlers"
	"github.com/vvoloshin/link-shortener/server"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct {
	Driver    string
	Port      string
	ShortBase string
	DBFile    string
	DBDir     string
}

func main() {
	config := readConfig()
	initDb(config)
	sqliteServer := server.NewServer(config.Port, config.DBFile, config.Driver)
	err := sqliteServer.Storage.InitTables()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/encode", handlers.EncodeUrl(config.ShortBase, sqliteServer.Storage))
	http.Handle("/bundle", handlers.BundleUrl(config.ShortBase, sqliteServer.Storage))
	http.Handle("/decode", handlers.DecodeUrl(sqliteServer.Storage))
	http.Handle("/redirect/", handlers.Redirect("/redirect/", sqliteServer.Storage))
	log.Println("starts server at port: " + sqliteServer.Port)
	err = http.ListenAndServe(sqliteServer.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initDb(config *Config) {
	if isFileNotExist(config.DBFile) {
		log.Println("database empty, created it")
		createFile(config.DBDir, config.DBFile)
	}
	log.Println("found existing database file")
}

func readConfig() *Config {
	return &Config{
		Driver:    "sqlite3",
		Port:      ":8080",
		ShortBase: "https://short.com",
		DBFile:    filepath.FromSlash("sqlite\\base.db"),
		DBDir:     filepath.FromSlash("sqlite"),
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
