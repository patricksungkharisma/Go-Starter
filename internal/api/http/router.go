package http

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/patricksungkharisma/go-starter/internal/middleware"
)

func initRouting(handler Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "POST, GET, PUT, DELETE",
		ValidateHeaders: false,
	}))

	router.GET("/example/:id", middleware.WrapHandlerFunc(handler.AppHandler.GetExample))

	return router
}
