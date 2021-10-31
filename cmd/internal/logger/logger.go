package logger

import (
	"github.com/ktakenaka/gosample2021/config"
	"github.com/ktakenaka/gosample2021/infrastructure/logger"
)

func Initialize(cfg config.Config) (logger.Logger, error) {
	return logger.New(cfg.Logger), nil
}
