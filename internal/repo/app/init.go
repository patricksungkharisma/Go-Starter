package app

import (
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	database databaseResource
}

func New(
	database *sqlx.DB,
) *Repo {
	return &Repo{
		database: database,
	}
}
