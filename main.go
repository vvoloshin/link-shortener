package main

import (
	"link-shortener/handlers"
	"link-shortener/storage"
	"log"
	"net/http"
	"os"
)

const port = "8080"

func main() {

	valueMap := new(storage.InMemStorage)
	valueMap.Store = map[string]string{}

	http.Handle("/encode", handlers.EncodeUrl(valueMap))
	http.Handle("/decode", handlers.DecodeUrl(valueMap))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}
