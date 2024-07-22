package app

import "github.com/patricksungkharisma/go-starter/internal/config"

type Handler struct {
	cfg        config.Config
	appUsecase appUsecase
}

func New(
	config config.Config,
	appUsecase appUsecase,
) *Handler {
	return &Handler{
		cfg:        config,
		appUsecase: appUsecase,
	}
}
