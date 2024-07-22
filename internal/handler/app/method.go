package app

import (
	"context"
	"log"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
	entityerror "github.com/patricksungkharisma/go-starter/internal/error"
	"github.com/patricksungkharisma/go-starter/internal/middleware"
)

func (h *Handler) GetExample(r *middleware.RequestContext) {
	ctx := context.Background()
	id := r.GinContext.Param("id")

	request := entity.GetExampleRequest{
		ID: id,
	}

	result, err := h.appUsecase.GetExampleData(ctx, request)
	if err != nil {
		log.Println(err)
		middleware.ResponseFail(r, middleware.HTTPServerErrorStatusCode, entityerror.GetExampleDataError, err)
	}

	middleware.ResponseSuccess(r, result)
}
