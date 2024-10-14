package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/patricksungkharisma/go-starter/internal/config"
)

type Repo struct {
	Config   config.Config
	Database DatabaseResource
}

func New(
	config config.Config,
	database *sqlx.DB,
) *Repo {
	return &Repo{
		Config:   config,
		Database: database,
	}
}
