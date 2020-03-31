package storage

type Storage interface {
	Read(key string) (string, error)
	Save(key string, value string) error
	InitTable() error
}
