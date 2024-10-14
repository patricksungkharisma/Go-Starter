package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/patricksungkharisma/go-starter/internal/error"
)

func WrapHandlerFunc(funcHandler func(r *RequestContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &RequestContext{
			GinContext: c,
		}
		funcHandler(ctx)
	}
}

func ResponseSuccess(ctx *RequestContext, data interface{}) {
	c := ctx.GinContext
	if data == nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	response := SuccessResponse{
		Response: data,
	}

	c.JSON(http.StatusOK, response)
}

func ResponseFail(ctx *RequestContext, message string, err *errs.TemplateError) {
	c := ctx.GinContext

	errors := ErrorResponse{
		Message:      message,
		ErrorDetails: err.Error,
	}

	c.JSON(err.StatusCode, errors)
}
