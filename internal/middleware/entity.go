package middleware

import "github.com/gin-gonic/gin"

const (
	HTTPSuccessStatusCode     = 200
	HTTPServerErrorStatusCode = 500
)

type RequestContext struct {
	GinContext *gin.Context
}

type Response struct {
	Response interface{} `json:"response"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
