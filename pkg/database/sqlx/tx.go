package sqlx

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

//go:generate mockery --name=TxProvider --outpkg=mock --output=./mock
type TxProvider interface {
	CloseTx(tx *sqlx.Tx, err error) error
	BeginTx(ctx context.Context) (*sqlx.Tx, error)
}

type Tx struct {
	db DBer
}

func NewTx(db DBer) *Tx {
	return &Tx{db: db}
}

func (t *Tx) BeginTx(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create new transaction")
		return nil, pkgerr.ErrWithStackTrace(err)
	}
	return tx, nil
}

func (t *Tx) CloseTx(tx *sqlx.Tx, err error) error {
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			log.Error().Err(err).Msg("failed to rollback transaction")
			return pkgerr.ErrWithStackTrace(err)
		}
		return nil
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msg("failed to commit transaction")
		return pkgerr.ErrWithStackTrace(err)
	}
	return nil
}
