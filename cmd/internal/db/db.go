package db

import (
	"github.com/ktakenaka/gosample2021/config"

	"github.com/ktakenaka/gosample2021/ent"
	"github.com/ktakenaka/gosample2021/infrastructure/database"
)

func Initialize(cfg config.Config) (*ent.Client, error) {
	return database.New(cfg.DB)
}
