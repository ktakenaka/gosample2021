package config

import (
	"errors"
	"os"

	cfg "github.com/ktakenaka/gosample/config"
)

func Initialize() (cfg.Config, error) {
	var cfgPath string
	switch os.Getenv("ENV") {
	case "development":
		cfgPath = "environment/development/config.yml"
	default:
		return cfg.Config{}, errors.New("not having configuration")
	}
	return cfg.New(cfgPath)
}
