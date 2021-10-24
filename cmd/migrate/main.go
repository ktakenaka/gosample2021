package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ktakenaka/gosample/cmd/internal/config"
	"github.com/ktakenaka/gosample/cmd/internal/db"
	"github.com/ktakenaka/gosample/ent/migrate"
)

func main() {
	cfg, err := config.Initialize()
	if err != nil {
		log.Fatalf("failed initializing cofig: %v", err)
	}

	client, err := db.Initialize(cfg)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
