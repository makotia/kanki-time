package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var conf Config

// Config is config struct
type Config struct {
	Server struct {
		StaticDir string `toml:"StaticDir"`
		TTFPath   string `toml:"TTFPath"`
	} `toml:"Server"`
}

func init() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
}

// GetConfig is return front directory
func GetConfig() Config {
	return conf
}
