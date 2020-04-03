package storage

import (
	"fmt"
	"github.com/vvoloshin/link-shortener/dbmodels"
	"log"
	"time"
)

func (s SQLite) Read(key string) (string, error) {
	db := s.DB
	rows, err := db.Query("SELECT * FROM URLS WHERE HASHED = $1", key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	m := dbmodels.UrlModel{}
	var timestamp int64
	for rows.Next() {
		err := rows.Scan(&m.Hashed, &m.Url, &timestamp)
		if err != nil {
			log.Fatal(err)
		}
	}
	m.Created = time.Unix(timestamp, 0)
	if m.Url == "" {
		return "", fmt.Errorf("not found url by key")
	}
	return m.Url, nil
}

func (s SQLite) Save(key string, value string) {
	db := s.DB
	_, err := db.Exec("INSERT INTO URLS (HASHED, URL, CREATED) VALUES ($1, $2, $3)", key, value, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inserted row with key: ", key)
}

func (s SQLite) Archive(key string) error {
	panic("not implemented")
}

func (s SQLite) InitTables() error {
	db := s.DB
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS URLS (HASHED TEXT PRIMARY KEY NOT NULL, URL TEXT NOT NULL,CREATED NUMBER NOT NULL)")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS URLS_ARCHIVED (HASHED TEXT PRIMARY KEY NOT NULL, URL TEXT NOT NULL,CREATED NUMBER NOT NULL, ARCHIVED TEXT NOT NULL)")
	if err != nil {
		return err
	}
	return nil
}
