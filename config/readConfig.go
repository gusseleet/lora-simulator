package config

import (
	"github.com/BurntSushi/toml"
	"os"
)



func ReadConfig(v interface {}) (interface {}, error) {
	var configfile = "test.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		return nil, err
	}

	if _, err := toml.DecodeFile(configfile, &v); err != nil {
		return nil, err
	}

	return v, nil
}