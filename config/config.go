package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ktakenaka/gosample/infrastructure/database"
)

type Config struct {
	DB *database.Config
}

func New(configFilePath string) (Config, error) {
	cfg := Config{}
	cfgByte, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	cfgByte = []byte(os.ExpandEnv(string(cfgByte)))
	err = yaml.Unmarshal(cfgByte, &cfg)

	cfgByte = []byte(os.ExpandEnv(string(cfgByte)))
	err = yaml.Unmarshal(cfgByte, &cfg)
	return cfg, err
}
