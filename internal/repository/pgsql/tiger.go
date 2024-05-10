package pgsql

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/helper/pagination"
	"github.com/vekaputra/tiger-kittens/internal/model"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	pkgsqlx "github.com/vekaputra/tiger-kittens/pkg/database/sqlx"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

const (
	TigerTable         = "tigers"
	TigerSightingTable = "tiger_sightings"
)

//go:generate mockery --name=TigerRepositoryProvider --outpkg=mock --output=../mock
type TigerRepositoryProvider interface {
	pkgsqlx.TxProvider
	Count(ctx context.Context) (uint64, error)
	CountSighting(ctx context.Context) (uint64, error)
	FindByName(ctx context.Context, name string) ([]entity.Tiger, error)
	FindWithPagination(ctx context.Context, page model.PaginationRequest) ([]entity.Tiger, error)
	FindSightingWithPagination(ctx context.Context, page model.PaginationRequest) ([]entity.Sighting, error)
	Insert(ctx context.Context, entity entity.Tiger) error
	TxFindByID(ctx context.Context, tx *sqlx.Tx, id int) (*entity.Tiger, error)
	TxInsertSighting(ctx context.Context, tx *sqlx.Tx, entity entity.TigerSighting) error
	TxUpdate(ctx context.Context, tx *sqlx.Tx, entity entity.Tiger) error
}

type TigerRepository struct {
	*pkgsqlx.Tx
	db pkgsqlx.DBer
	sb squirrel.StatementBuilderType
}

func NewTigerRepository(db pkgsqlx.DBer) *TigerRepository {
	return &TigerRepository{
		Tx: pkgsqlx.NewTx(db),
		db: db,
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *TigerRepository) FindByName(ctx context.Context, name string) ([]entity.Tiger, error) {
	query, args, err := r.sb.Select(
		"id",
		"date_of_birth",
		"last_lat",
		"last_long",
		"last_photo",
		"last_seen",
		"name",
		"created_at",
		"updated_at",
	).
		From(TigerTable).
		Where(squirrel.Eq{"name": name}).
		ToSql()
	if err != nil {
		return []entity.Tiger{}, pkgerr.ErrWithStackTrace(err)
	}

	var result []entity.Tiger
	if err = r.db.SelectContext(ctx, &result, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to find tigers by name")
		return []entity.Tiger{}, pkgerr.ErrWithStackTrace(err)
	}

	return result, nil
}

func (r *TigerRepository) Insert(ctx context.Context, entity entity.Tiger) error {
	query, args, err := r.sb.Insert(TigerTable).
		Columns("date_of_birth", "last_lat", "last_long", "last_photo", "last_seen", "name").
		Values(
			entity.DateOfBirth.Format(time.DateOnly),
			entity.LastLat,
			entity.LastLong,
			entity.LastPhoto,
			entity.LastSeen.Format(time.RFC3339),
			entity.Name,
		).
		ToSql()
	if err != nil {
		return pkgerr.ErrWithStackTrace(err)
	}

	if _, err = r.db.ExecContext(ctx, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to insert new tiger")
		return pkgerr.ErrWithStackTrace(err)
	}

	return nil
}

func (r *TigerRepository) Count(ctx context.Context) (uint64, error) {
	query, args, err := r.sb.Select("COUNT(1)").
		From(TigerTable).
		ToSql()
	if err != nil {
		return 0, pkgerr.ErrWithStackTrace(err)
	}

	var totalItem uint64
	if err = r.db.GetContext(ctx, &totalItem, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to count tigers")
		return 0, pkgerr.ErrWithStackTrace(err)
	}

	return totalItem, nil
}

func (r *TigerRepository) FindWithPagination(ctx context.Context, page model.PaginationRequest) ([]entity.Tiger, error) {
	query, args, err := r.sb.Select(
		"id",
		"date_of_birth",
		"last_lat",
		"last_long",
		"last_photo",
		"last_seen",
		"name",
		"created_at",
		"updated_at",
	).
		From(TigerTable).
		OrderBy("last_seen DESC").
		Offset(pagination.SQLOffset(page.Page, page.PerPage)).
		Limit(page.PerPage).
		ToSql()
	if err != nil {
		return []entity.Tiger{}, pkgerr.ErrWithStackTrace(err)
	}

	var result []entity.Tiger
	if err = r.db.SelectContext(ctx, &result, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to find tigers with pagination")
		return []entity.Tiger{}, pkgerr.ErrWithStackTrace(err)
	}

	return result, nil
}

func (r *TigerRepository) TxFindByID(ctx context.Context, tx *sqlx.Tx, id int) (*entity.Tiger, error) {
	query, args, err := r.sb.Select(
		"id",
		"date_of_birth",
		"last_lat",
		"last_long",
		"last_photo",
		"last_seen",
		"name",
		"created_at",
		"updated_at",
	).
		From(TigerTable).
		Where(squirrel.Eq{"id": id}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, pkgerr.ErrWithStackTrace(err)
	}

	var result entity.Tiger
	if err = tx.GetContext(ctx, &result, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to find tigers by id")
		return nil, pkgerr.ErrWithStackTrace(err)
	}

	return &result, nil
}

func (r *TigerRepository) TxInsertSighting(ctx context.Context, tx *sqlx.Tx, entity entity.TigerSighting) error {
	query, args, err := r.sb.Insert(TigerSightingTable).
		Columns("user_id", "tiger_id", "lat", "long", "photo").
		Values(
			entity.UserID,
			entity.TigerID,
			entity.Lat,
			entity.Long,
			entity.Photo,
		).
		ToSql()
	if err != nil {
		return pkgerr.ErrWithStackTrace(err)
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to insert new tiger sighting")
		return pkgerr.ErrWithStackTrace(err)
	}

	return nil
}

func (r *TigerRepository) TxUpdate(ctx context.Context, tx *sqlx.Tx, entity entity.Tiger) error {
	query, args, err := r.sb.Update(TigerTable).
		Set("last_lat", entity.LastLat).
		Set("last_long", entity.LastLong).
		Set("last_photo", entity.LastPhoto).
		Set("last_seen", entity.LastSeen).
		Where(squirrel.Eq{"id": entity.ID}).
		ToSql()
	if err != nil {
		return pkgerr.ErrWithStackTrace(err)
	}

	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to update tiger")
		return pkgerr.ErrWithStackTrace(err)
	}

	return nil
}

func (r *TigerRepository) FindSightingWithPagination(ctx context.Context, page model.PaginationRequest) ([]entity.Sighting, error) {
	query, args, err := r.sb.Select(
		"u.username as username",
		"t.name as tiger_name",
		"ts.photo",
		"ts.lat",
		"ts.long",
		"ts.created_at",
	).
		From(fmt.Sprintf(`%s ts`, TigerSightingTable)).
		Join(fmt.Sprintf(`%s u ON u.id = ts.user_id`, UserTable)).
		Join(fmt.Sprintf(`%s t ON t.id = ts.tiger_id`, TigerTable)).
		OrderBy("created_at DESC").
		Offset(pagination.SQLOffset(page.Page, page.PerPage)).
		Limit(page.PerPage).
		ToSql()
	if err != nil {
		return []entity.Sighting{}, pkgerr.ErrWithStackTrace(err)
	}

	var result []entity.Sighting
	if err = r.db.SelectContext(ctx, &result, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to find sightings with pagination")
		return []entity.Sighting{}, pkgerr.ErrWithStackTrace(err)
	}

	return result, nil
}

func (r *TigerRepository) CountSighting(ctx context.Context) (uint64, error) {
	query, args, err := r.sb.Select("COUNT(1)").
		From(TigerSightingTable).
		ToSql()
	if err != nil {
		return 0, pkgerr.ErrWithStackTrace(err)
	}

	var totalItem uint64
	if err = r.db.GetContext(ctx, &totalItem, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to count sightings")
		return 0, pkgerr.ErrWithStackTrace(err)
	}

	return totalItem, nil
}
