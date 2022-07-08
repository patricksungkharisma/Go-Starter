package main

import (
	"github.com/gin-gonic/gin"

	cors "github.com/itsjamie/gin-cors"
	appHandler "github.com/patricksungkharisma/go-starter/internal/handler/app"
)

func httpStarter(handler *Handler) error {

	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.Middleware(cors.Config{
		Origins: "*",
		Methods: "POST, GET, PUT, DELETE",
		ValidateHeaders: false,
	}))

	router.GET("/", appHandler.WrapHandlerFunc(appHandler.GetHelloWorld, handler.appHandler))

	// Listen and Serve to This Port
	router.Run(":8080")

	return nil
}