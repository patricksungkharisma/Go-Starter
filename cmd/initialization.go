package cmd

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/patricksungkharisma/go-starter/internal/config"
	apphandler "github.com/patricksungkharisma/go-starter/internal/handler/app"
	apprepo "github.com/patricksungkharisma/go-starter/internal/repo/app"
	appusecase "github.com/patricksungkharisma/go-starter/internal/usecase/app"
	"github.com/pkg/errors"
)

func InitResource(cfg config.Config) (Resource, error) {
	resource := Resource{}

	connString := fmt.Sprintf(cfg.Database.DatabaseConnectionFormat, cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DatabaseName)
	database, err := sqlx.Connect(cfg.Database.DatabaseType, connString)
	if err != nil {
		return resource, errors.Wrap(err, "[initResource]")
	}

	resource.Database = database

	return resource, nil
}

func InitRepo(cfg config.Config, resource Resource) Repo {
	repo := Repo{}

	app := apprepo.New(cfg, resource.Database)
	repo.AppRepo = app

	return repo
}

func InitUsecase(cfg config.Config, repo Repo) Usecase {
	usecase := Usecase{}

	app := appusecase.New(cfg, repo.AppRepo)
	usecase.AppUsecase = app

	return usecase
}

func InitHandler(cfg config.Config, usecase Usecase) Handler {
	handler := Handler{}

	app := apphandler.New(cfg, usecase.AppUsecase)
	handler.AppHandler = app

	return handler
}
