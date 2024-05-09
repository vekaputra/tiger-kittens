package pgsql

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	pkgsqlx "github.com/vekaputra/tiger-kittens/pkg/database/sqlx"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

const (
	UserTable = "users"
)

//go:generate mockery --name=UserRepositoryProvider --outpkg=mock --output=../mock
type UserRepositoryProvider interface {
	FindByEmailOrUsername(ctx context.Context, email, username string) ([]entity.User, error)
	Insert(ctx context.Context, entity entity.User) error
}

type UserRepository struct {
	db pkgsqlx.DBer
	sb squirrel.StatementBuilderType
}

func NewUserRepository(db pkgsqlx.DBer) *UserRepository {
	return &UserRepository{
		db: db,
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *UserRepository) FindByEmailOrUsername(ctx context.Context, email, username string) ([]entity.User, error) {
	query, args, err := r.sb.Select("id", "email", "password", "username", "created_at").
		From(UserTable).
		Where(squirrel.Or{
			squirrel.Eq{"email": email},
			squirrel.Eq{"username": username},
		}).
		ToSql()
	if err != nil {
		return []entity.User{}, pkgerr.ErrWithStackTrace(err)
	}

	var result []entity.User
	if err = r.db.SelectContext(ctx, &result, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to find by email or username")
		return []entity.User{}, pkgerr.ErrWithStackTrace(err)
	}

	return result, nil
}

func (r *UserRepository) Insert(ctx context.Context, entity entity.User) error {
	query, args, err := r.sb.Insert(UserTable).
		Columns("email", "password", "username").
		Values(entity.Email, entity.Password, entity.Username).
		ToSql()
	if err != nil {
		return pkgerr.ErrWithStackTrace(err)
	}

	if _, err = r.db.ExecContext(ctx, query, args...); err != nil {
		log.Error().Err(err).Msg("failed to insert new user")
		return pkgerr.ErrWithStackTrace(err)
	}

	return nil
}
