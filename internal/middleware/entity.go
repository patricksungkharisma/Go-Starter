package middleware

import "github.com/gin-gonic/gin"

type RequestContext struct {
	GinContext *gin.Context
}

type SuccessResponse struct {
	Response interface{} `json:"response"`
}

type ErrorResponse struct {
	Message      string `json:"message"`
	ErrorDetails error  `json:"error_details,omitempty"`
}
