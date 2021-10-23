package main

import (
    "context"
    "log"

	_ "github.com/go-sql-driver/mysql"
    
    "github.com/ktakenaka/gosample/ent"
    "github.com/ktakenaka/gosample/ent/migrate"
)

func main() {
	// ここで条件分岐して、環境ごとにべつのconfigでできるようにする
    client, err := ent.Open("mysql", "gosample:password@tcp(database:3306)/gosample_development")
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
