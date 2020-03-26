package main

import (
	"github.com/vvoloshin/link-shortener/handlers"
	"github.com/vvoloshin/link-shortener/storage"
	"log"
	"net/http"
)

const port = "8080"

func main() {

	valueMap := new(storage.InMemStorage)
	valueMap.Store = map[string]string{}

	http.Handle("/encode", handlers.EncodeUrl(valueMap))
	http.Handle("/decode", handlers.DecodeUrl(valueMap))

	log.Println("starts server at port: " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
