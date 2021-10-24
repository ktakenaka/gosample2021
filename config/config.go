package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ktakenaka/gosample/infrastructure/database"
	"github.com/ktakenaka/gosample/infrastructure/logger"
)

type Config struct {
	DB     *database.Config
	Logger *logger.Config
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
