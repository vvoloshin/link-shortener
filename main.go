package main

import (
	"github.com/vvoloshin/link-shortener/config"
	"github.com/vvoloshin/link-shortener/handlers"
	"log"
	"net/http"
)

func main() {

	s := config.NewDefault()
	http.Handle("/encode", handlers.EncodeUrl(s.Storage))
	http.Handle("/decode", handlers.DecodeUrl(s.Storage))
	http.Handle("/redirect", handlers.Redirect(s.Storage))

	log.Println("starts server at port: " + s.Port)
	err := http.ListenAndServe(s.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
