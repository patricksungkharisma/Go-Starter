package main

import (
	appRepo "github.com/patricksungkharisma/go-starter/internal/repo/app"

	appUsecase "github.com/patricksungkharisma/go-starter/internal/usecase/app"

	appHandler "github.com/patricksungkharisma/go-starter/internal/handler/app"
)

type resource struct {
}

type repo struct {
	appRepo *appRepo.Repo
}

type usecase struct {
	appUsecase *appUsecase.Usecase
}

type Handler struct {
	appHandler *appHandler.Handler
}