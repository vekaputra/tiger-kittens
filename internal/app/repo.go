package app

import "github.com/vekaputra/tiger-kittens/internal/repository/pgsql"

type Repo struct {
	TigerRepo pgsql.TigerRepositoryProvider
	UserRepo  pgsql.UserRepositoryProvider
}

func NewRepo(conn *Connection) Repo {
	return Repo{
		TigerRepo: pgsql.NewTigerRepository(conn.DB),
		UserRepo:  pgsql.NewUserRepository(conn.DB),
	}
}
