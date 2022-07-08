package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HandlerRequest struct {
	GinContext *gin.Context
	handler *Handler
}

func WrapHandlerFunc(funcHandler func(r *HandlerRequest), handler *Handler) gin.HandlerFunc {

	return func(c *gin.Context) {

		r := &HandlerRequest{
			GinContext: c,
			handler: handler,
		}

		funcHandler(r)
	}
	
}

func GetHelloWorld(r *HandlerRequest) {

	fmt.Println("Hello World")

	return
}