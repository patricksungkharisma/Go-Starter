package main

import (
	"log"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/patricksungkharisma/go-starter/cmd"
	"github.com/patricksungkharisma/go-starter/internal/config"
	"github.com/patricksungkharisma/go-starter/internal/middleware"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalln("[main] error on init config: ", err)
	}

	resource, err := cmd.InitResource(config)
	if err != nil {
		log.Fatalln("[main] error on init resource: ", err)
	}

	repo := cmd.InitRepo(config, resource)
	usecase := cmd.InitUsecase(config, repo)
	handler := cmd.InitHandler(config, usecase)

	err = serverStarter(handler)
	if err != nil {
		log.Fatalln("[main] error when starting server: ", err)
	}
}

func serverStarter(handler cmd.Handler) error {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "POST, GET, PUT, DELETE",
		ValidateHeaders: false,
	}))

	// routing
	router.GET("/example/:id", middleware.WrapHandlerFunc(handler.AppHandler.GetExample))

	router.Run(":8080")

	return nil
}
