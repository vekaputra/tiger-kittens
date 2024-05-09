package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/vekaputra/tiger-kittens/internal/config"
	pkgsqlx "github.com/vekaputra/tiger-kittens/pkg/database/sqlx"
)

func NewDatabase(config config.DatabaseConfig) (pkgsqlx.DBer, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.Username, config.Password, config.DBName)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(config.MaxIdleConnection)
	db.SetMaxOpenConns(config.MaxOpenConnection)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
