package app

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
	errs "github.com/patricksungkharisma/go-starter/internal/error"
)

func (u *Usecase) GetExampleData(ctx context.Context, req entity.GetExampleRequest) (*entity.GetExampleResponse, *errs.TemplateError) {
	parsedIDInt, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		slog.ErrorContext(ctx, errs.ErrTagGetExampleData,
			err,
			slog.Any("id", req.ID),
		)
		return nil, errs.GenerateError(http.StatusInternalServerError, err)
	}

	result, customErr := u.AppRepo.GetExampleDataDB(ctx, parsedIDInt)
	if customErr != nil {
		return nil, customErr
	}

	return result, nil
}
