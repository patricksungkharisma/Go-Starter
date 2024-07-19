package app

import (
	"context"
	"log"
	"strconv"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
	entityerror "github.com/patricksungkharisma/go-starter/internal/error"
	"github.com/patricksungkharisma/go-starter/internal/middleware"
)

func (h *Handler) GetExample(r *middleware.RequestContext) {
	ctx := context.Background()
	id := r.GinContext.Param("id")
	parsedIDInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
		middleware.ResponseFail(r, middleware.HTTPServerErrorStatusCode, entityerror.ParameterParsingError, err)
		return
	}

	request := entity.GetExampleRequest{
		ID: parsedIDInt,
	}

	result, err := h.AppUsecase.GetExampleData(ctx, request)
	if err != nil {
		log.Println(err)
		middleware.ResponseFail(r, middleware.HTTPServerErrorStatusCode, entityerror.GetExampleDataError, err)
	}

	middleware.ResponseSuccess(r, result)
}
