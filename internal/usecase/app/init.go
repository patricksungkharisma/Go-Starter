package app

import "github.com/patricksungkharisma/go-starter/internal/config"

type Usecase struct {
	cfg     config.Config
	appRepo appRepo
}

func New(
	config config.Config,
	appRepo appRepo,
) *Usecase {
	return &Usecase{
		cfg:     config,
		appRepo: appRepo,
	}
}
