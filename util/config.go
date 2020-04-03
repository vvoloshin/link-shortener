package util

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
)

type ApiSecure struct {
	ApiKey    string
	ApiHeader string
}

type DBConfig struct {
	Driver string
	DBFile string
	DBDir  string
}

type ServerHost struct {
	Port string
	Host string
}

type Config struct {
	ServerHost ServerHost
	ApiSecure  ApiSecure
	DBConfig   DBConfig
}

func ReadConfig() (*Config, error) {
	var conf Config
	path := filepath.FromSlash("config/properties.toml")
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, err
	} else {
		return &conf, nil
	}
}
