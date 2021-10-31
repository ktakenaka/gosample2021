package mysql

import (
	"context"
	"fmt"

	"github.com/ktakenaka/gosample2021/config"
	"github.com/ktakenaka/gosample2021/ent"
	"github.com/ktakenaka/gosample2021/infrastructure/database"
)

func GetClient(cfg config.Config) (client *ent.Client, release func()) {
	// Put `map[string]string{"foreign_key_check": "0"}` to cfg.DB.Options
	client, err := database.New(cfg.DB)
	if err != nil {
		panic(err)
	}

	release = func() {
		if err := cleanDB(client); err != nil {
			panic(err)
		}
		if err := client.Close(); err != nil {
			panic(err)
		}
	}
	return
}

func cleanDB(client *ent.Client) (err error) {
	// TODO: get table names from INFORMATION_SCHEMA so that we don't need to define tables
	// "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA='gosample2021_test' AND TABLE_NAME != 'schema_migrations'"
	ctx := context.TODO()
	tx, err := client.Tx(ctx)
	if err != nil {
		panic(err)
	}

	_, err = tx.Sample.Delete().Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	_, err = client.Office.Delete().Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
