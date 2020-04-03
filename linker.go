package main

import (
	"github.com/vvoloshin/link-shortener/handlers"
	"github.com/vvoloshin/link-shortener/server"
	"github.com/vvoloshin/link-shortener/util"
	"log"
	"net/http"
	"os"
)

func main() {
	config := util.ReadConfig()
	initDb(config)
	sqliteServer := server.NewServer(config.ServerHost.Port, config.DBConfig.DBFile, config.DBConfig.Driver)
	err := sqliteServer.Storage.InitTables()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", handlers.EncodeUrl(config, sqliteServer.Storage))
	http.Handle("/bundle", handlers.BundleUrl(config, sqliteServer.Storage))
	http.Handle("/", handlers.Redirect(config, sqliteServer.Storage))
	log.Println("starts server at port " + sqliteServer.Port)
	err = http.ListenAndServe(sqliteServer.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initDb(config *util.Config) {
	if isFileNotExist(config.DBConfig.DBFile) {
		log.Println("database empty, try to create it")
		createFile(config.DBConfig.DBDir, config.DBConfig.DBFile)
		return
	}
	log.Println("found existing database file")
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
