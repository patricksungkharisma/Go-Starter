package app

import (
	"context"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
	errs "github.com/patricksungkharisma/go-starter/internal/error"
	"github.com/patricksungkharisma/go-starter/internal/middleware"
)

func (h *Handler) GetExample(r *middleware.RequestContext) {
	ctx := context.Background()
	id := r.GinContext.Param("id")

	request := entity.GetExampleRequest{
		ID: id,
	}

	result, customErr := h.AppUsecase.GetExampleData(ctx, request)
	if customErr != nil {
		middleware.ResponseFail(r, errs.ErrGetExampleData, customErr)
	}

	middleware.ResponseSuccess(r, result)
}
