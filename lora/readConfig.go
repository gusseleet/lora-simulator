package lora

import (
	"github.com/BurntSushi/toml"
	"os"
	"lora_simulator/logging"
	 "github.com/sanity-io/litter"
)



func ReadConfig(v LoraConfig, path string) (LoraConfig, error) {
	_, err := os.Stat(path)

	if err != nil {
		return v, err
	}

	if _, err := toml.DecodeFile(path, &v); err != nil {
		return v, err
	}

	logging.Log().Debug("Config: ", litter.Sdump(v))

	return v, nil
}