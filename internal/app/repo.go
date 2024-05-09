package app

import "github.com/vekaputra/tiger-kittens/internal/repository/pgsql"

type Repo struct {
	UserRepo pgsql.UserRepositoryProvider
}

func NewRepo(conn *Connection) Repo {
	return Repo{
		UserRepo: pgsql.NewUserRepository(conn.DB),
	}
}
