package app

import "github.com/patricksungkharisma/go-starter/internal/config"

type Usecase struct {
	Config  config.Config
	AppRepo AppRepo
}

func New(
	config config.Config,
	appRepo AppRepo,
) *Usecase {
	return &Usecase{
		Config:  config,
		AppRepo: appRepo,
	}
}
