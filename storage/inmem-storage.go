package storage

import "fmt"

type InMemStorage struct {
	Store map[string]string
}

func (storage InMemStorage) Read(key string) (string, error) {
	if val, err := storage.Store[key]; err {
		return val, nil
	}
	return "", fmt.Errorf("not found value: %s", key)
}

func (storage InMemStorage) Save(key string, value string) error {
	storage.Store[key] = value
	return nil
}
