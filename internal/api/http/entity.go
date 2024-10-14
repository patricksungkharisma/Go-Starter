package http

import (
	apphandler "github.com/patricksungkharisma/go-starter/internal/handler/app"
	apprepo "github.com/patricksungkharisma/go-starter/internal/repo/app"
	appusecase "github.com/patricksungkharisma/go-starter/internal/usecase/app"
)

type Repo struct {
	AppRepo *apprepo.Repo
}

type Usecase struct {
	AppUsecase *appusecase.Usecase
}

type Handler struct {
	AppHandler *apphandler.Handler
}
