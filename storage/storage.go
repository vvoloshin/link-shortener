package storage

type Storage interface {
	Read(key string) (string, error)
	Archive(key string) error
	Save(key string, value string) error
	InitTables() error
}
