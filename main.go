package main

import (
	"github.com/BurntSushi/toml"
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
	http.Handle("/redirect/", handlers.Redirect("/redirect/", sqliteServer.Storage))
	log.Println("starts server at port: " + sqliteServer.Port)
	err = http.ListenAndServe(sqliteServer.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initDb(config *Config) {
	if isFileNotExist(config.DBFile) {
		log.Println("database empty, try to create it")
		createFile(config.DBDir, config.DBFile)
		return
	}
	log.Println("found existing database file")
}

func readConfig() *Config {
	var conf Config
	path := filepath.FromSlash("config\\properties.toml")
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Fatal("can't read configuration file from path=", path)
		return nil
	} else {
		return &conf
	}
}

func createFile(d, f string) {
	err := os.Mkdir(d, 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(f)
	if err != nil {
		log.Fatal(err)
	}
}

func isFileNotExist(f string) bool {
	_, err := os.Stat(f)
	return os.IsNotExist(err)
}
