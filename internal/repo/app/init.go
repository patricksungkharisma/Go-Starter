package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/patricksungkharisma/go-starter/internal/config"
)

type Repo struct {
	cfg      config.Config
	database databaseResource
}

func New(
	config config.Config,
	database *sqlx.DB,
) *Repo {
	return &Repo{
		cfg:      config,
		database: database,
	}
}
