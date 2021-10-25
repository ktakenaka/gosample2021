package config

import (
	"errors"
	"flag"

	conf "github.com/ktakenaka/gosample/config"
)

func Initialize() (conf.Config, error) {
	configFilePath := flag.String("c", "", "config file path for app")
	flag.Parse()
	if configFilePath == nil {
		return conf.Config{}, errors.New("Not having config file")
	}

	return conf.New(*configFilePath)
}
