package app

import (
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/config"
	pkgsqlx "github.com/vekaputra/tiger-kittens/pkg/database/sqlx"
)

type Connection struct {
	DB pkgsqlx.DBer
}

func NewConnection(config *config.Config) *Connection {
	db, err := NewDatabase(config.DatabaseConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed when initiate database")
	}

	return &Connection{
		DB: db,
	}
}
