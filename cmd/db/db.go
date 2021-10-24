package db

import (
	"github.com/ktakenaka/gosample/config"

	"github.com/ktakenaka/gosample/ent"
	"github.com/ktakenaka/gosample/infrastructure/database"
)

func Initialize(cfg config.Config) (*ent.Client, error) {
	return database.New(cfg.DB)
}
