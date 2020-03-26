package storage

import "fmt"

type InMemStorage struct {
	Store map[string]string
}

func (i InMemStorage) Read(key string) (string, error) {
	if val, err := i.Store[key]; err {
		return val, nil
	}
	return "", fmt.Errorf("not found value: %s", key)
}

func (i InMemStorage) Save(key string, value string) error {
	i.Store[key] = value
	return nil
}
