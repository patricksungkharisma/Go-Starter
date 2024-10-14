package app

import "github.com/patricksungkharisma/go-starter/internal/config"

type Handler struct {
	Config     config.Config
	AppUsecase AppUsecase
}

func New(
	config config.Config,
	appUsecase AppUsecase,
) *Handler {
	return &Handler{
		Config:     config,
		AppUsecase: appUsecase,
	}
}
