package mysql

import (
	"context"

	"github.com/ktakenaka/gosample/config"
	"github.com/ktakenaka/gosample/ent"
	"github.com/ktakenaka/gosample/infrastructure/database"
)

const ()

func GetDB() (client *ent.Client, release func()) {
	cfg, err := config.New("environment/test/config.yml")
	if err != nil {
		panic(err)
	}

	cfg.DB.Options = map[string]string{"foreign_key_check": "0"}
	client, err = database.New(cfg.DB)
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
	// "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA='gosample_test' AND TABLE_NAME != 'schema_migrations'"
	ctx := context.TODO()
	_, err = client.Sample.Delete().Exec(ctx)
	if err != nil {
		return err
	}

	_, err = client.Office.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
