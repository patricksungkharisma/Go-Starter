package main

import (
	"context"
	"fmt"

	appRepo "github.com/patricksungkharisma/go-starter/internal/repo/app"

	appUsecase "github.com/patricksungkharisma/go-starter/internal/usecase/app"

	appHandler "github.com/patricksungkharisma/go-starter/internal/handler/app"
)

func initApp() (error) {

	var (
		ctx = context.Background()
	)

	resource, errResource := initResource(ctx)
	if errResource != nil {
		fmt.Println("[app][initResource] Error when init resource")
		return errResource
	}

	repo, errRepo := initRepo(ctx, resource)
	if errRepo != nil {
		fmt.Println("[app][initRepo] Error when init repo")
		return errRepo
	}

	usecase, errUsecase := initUsecase(ctx, repo)
	if errUsecase != nil {
		fmt.Println("[app][initUsecase] Error when init usecase")
		return errUsecase
	}

	handler, errHandler := initHandler(ctx, usecase)
	if errHandler != nil {
		fmt.Println("[app][initUsecase] Error when init handler")
		return errHandler
	}

	// HTTP Server Starter
	errStart := httpStarter(handler)
	if errStart != nil {
		fmt.Println("[app][httpStarter] Error when starting HTTP server")
	}


	return nil
}

func initResource(ctx context.Context) (*resource, error) {

	resource := new(resource)

	return resource, nil
}

func initRepo(ctx context.Context, resource *resource) (*repo, error) {

	repo := new(repo)

	// App Repo
	app := appRepo.New()
	repo.appRepo = app

	return repo, nil
}

func initUsecase(ctx context.Context, repo *repo) (*usecase, error) {

	usecase := new(usecase)

	// App Usecase
	app := appUsecase.New(repo.appRepo)
	usecase.appUsecase = app

	return usecase, nil
}

func initHandler(ctx context.Context, usecase *usecase) (*Handler, error) {
	handler := new(Handler)

	// App Handler
	app := appHandler.New(usecase.appUsecase)
	handler.appHandler = app

	return handler, nil
}
