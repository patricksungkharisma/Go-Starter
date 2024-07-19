package middleware

import (
	"github.com/gin-gonic/gin"
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
		c.JSON(HTTPSuccessStatusCode, gin.H{})
		return
	}

	response := Response{
		Response: data,
	}

	c.JSON(HTTPSuccessStatusCode, response)
}

func ResponseFail(ctx *RequestContext, httpStatus int, message string, err error) {
	c := ctx.GinContext

	var errorDetails string
	if err != nil {
		errorDetails = err.Error()
	}

	errors := ErrorResponse{
		Message: message,
		Error:   errorDetails,
	}

	c.JSON(httpStatus, errors)
}
