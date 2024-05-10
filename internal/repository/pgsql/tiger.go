package pgsql

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/helper/pagination"
	"github.com/vekaputra/tiger-kittens/internal/model"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	pkgsqlx "github.com/vekaputra/tiger-kittens/pkg/database/sqlx"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

const (
	TigerTable = "tigers"
)

type TigerRepositoryProvider interface {
	Count(ctx context.Context) (uint64, error)
	FindByName(ctx context.Context, name string) ([]entity.Tiger, error)
	FindWithPagination(ctx context.Context, page model.PaginationRequest) ([]entity.Tiger, error)
	Insert(ctx context.Context, entity entity.Tiger) error
}

type TigerRepository struct {
	db pkgsqlx.DBer
	sb squirrel.StatementBuilderType
}

func NewTigerRepository(db pkgsqlx.DBer) *TigerRepository {
	return &TigerRepository{
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
