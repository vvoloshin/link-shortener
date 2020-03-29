package storage

import (
	"fmt"
	"github.com/vvoloshin/link-shortener/dbmodels"
	"log"
)

func (s SQLite) Read(key string) (string, error) {
	c := s.connect()
	defer c.Close()
	rows, err := c.Query("SELECT * FROM URLS WHERE HASHED = $1", key)
	if err != nil {
		log.Fatal("error with reading database file ")
	}

	m := dbmodels.UrlModel{}
	for rows.Next() {
		err := rows.Scan(&m.Hashed, &m.Url)
		if err != nil {
			return "", fmt.Errorf("not found url by key")
		}
	}
	if m.Url == "" {
		return "", fmt.Errorf("not found url by key")
	}
	return m.Url, nil
}

func (s SQLite) Save(key string, value string) error {
	c := s.connect()
	defer c.Close()
	_, err := c.Exec("INSERT INTO URLS (HASHED, URL) VALUES ($1, $2)", key, value)
	if err != nil {
		log.Println(err)
	}
	log.Println("inserted row with key: ", key)
	return nil
}
