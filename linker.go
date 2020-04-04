package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"

	"github.com/vvoloshin/link-shortener/handlers"
	"github.com/vvoloshin/link-shortener/server"
	"github.com/vvoloshin/link-shortener/util"
)

func main() {
	config, err := util.ReadConfig()
	util.CheckError(err)
	initDb(config)
	sqliteServer := server.NewServer("localhost:"+config.ServerHost.Port, config.DBConfig.DBFile, config.DBConfig.Driver)
	err = sqliteServer.Storage.InitTables()
	util.CheckError(err)
	http.Handle("/processing", handlers.Processing(config, sqliteServer.Storage))
	http.Handle("/", handlers.Redirect(sqliteServer.Storage))
	log.Println("starts server at port " + sqliteServer.Port)
	err = http.ListenAndServe(sqliteServer.Port, nil)
	util.CheckError(err)
}

func initDb(config *util.Config) {
	if isFileNotExist(config.DBConfig.DBFile) {
		log.Println("database file not exist, try to create it")
		createFile(config.DBConfig.DBDir, config.DBConfig.DBFile)
		return
	}
	log.Println("found existing database file")
}

func createFile(dir string, file string) {
	err := os.Mkdir(dir, 0755)
	_, err = os.Create(file)
	util.CheckError(err)
}

func isFileNotExist(file string) bool {
	_, err := os.Stat(file)
	return os.IsNotExist(err)
}
