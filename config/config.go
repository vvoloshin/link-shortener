package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"path/filepath"
)

type Config struct {
	Driver    string
	Port      string
	ShortBase string
	DBFile    string
	DBDir     string
	ApiKey    string
	ApiHeader string
}

func ReadConfig() *Config {
	var conf Config
	path := filepath.FromSlash("config\\properties.toml")
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Fatal("can't read configuration file from path=", path)
		return nil
	} else {
		return &conf
	}
}
