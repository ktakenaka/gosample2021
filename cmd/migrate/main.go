package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"

	"github.com/ktakenaka/gosample/ent/migrate"
	"github.com/ktakenaka/gosample/infrastructure/database"
)

func main() {
	cfg := database.Config{}
	if os.Getenv("ENV") == "development" {
		file, err := os.ReadFile("environment/development/config.yml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(file, &cfg)
		if err != nil {
			panic(err)
		}
	} else {
		panic("not having configuration")
	}

	client, err := database.New(&cfg)
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
