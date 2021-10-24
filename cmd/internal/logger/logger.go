package logger

import (
	"github.com/ktakenaka/gosample/config"
	"github.com/ktakenaka/gosample/infrastructure/logger"
)

func Initialize(cfg config.Config) (logger.Logger, error) {
	return logger.New(cfg.Logger), nil
}
