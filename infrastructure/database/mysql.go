package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/ktakenaka/gosample/ent"
)

// Config connection information
type Config struct {
	User            string
	Password        string
	Host            string
	DBName          string `yaml:"db_name"`
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration

	Options            map[string]string
	IsIgnoreForeignKey bool
	IsLogging          bool
}

const (
	conn          = "%s:%s@tcp(%s)/%s?collation=utf8mb4_bin&parseTime=true&charset=utf8mb4"
	connWithoutFK = conn + "&foreign_key_checks=0"
)

const (
	defaultMaxIdleConns    = 5
	defaultMaxOpenConns    = 10
	defaultConnMaxLifetime = 20 * time.Second
)

// New connect to db
func New(cfg *Config) (*ent.Client, error) {
	var connStr string
	if cfg.IsIgnoreForeignKey {
		connStr = connWithoutFK
	} else {
		connStr = conn
	}

	if cfg.Options != nil {
		var options []string
		for k, v := range cfg.Options {
			options = append(options, k+"="+v)
		}
		connStr = connStr + "&" + strings.Join(options, "&")
	}
	dst := fmt.Sprintf(connStr, cfg.User, cfg.Password, cfg.Host, cfg.DBName)

	db, err := connect(dst, cfg)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.MySQL, db)
	options := []ent.Option{ent.Driver(drv)}

	if cfg.IsLogging {
		options = append(options, ent.Debug())
	}

	return ent.NewClient(options...), nil
}

func connect(dst string, cfg *Config) (*sql.DB, error) {
	db, err := sql.Open(dialect.MySQL, dst)
	if err != nil {
		return nil, err
	}

	if cfg.MaxIdleConns == 0 {
		db.SetMaxIdleConns(defaultMaxIdleConns)
	} else {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	if cfg.MaxOpenConns == 0 {
		db.SetMaxIdleConns(defaultMaxOpenConns)
	} else {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	if cfg.ConnMaxLifetime == 0 {
		db.SetConnMaxLifetime(defaultConnMaxLifetime)
	} else {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}
	return db, nil
}
