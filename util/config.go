package util

import (
	"github.com/BurntSushi/toml"
	"log"
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
	Port      string
	ShortBase string
	Redirect  string
}

type Config struct {
	ServerHost
	ApiSecure
	DBConfig
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
