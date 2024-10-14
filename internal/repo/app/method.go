package app

import (
	"context"
	"log/slog"
	"net/http"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
	errs "github.com/patricksungkharisma/go-starter/internal/error"
)

func (r *Repo) GetExampleDataDB(ctx context.Context, id int64) (*entity.GetExampleResponse, *errs.TemplateError) {
	result := &entity.GetExampleResponse{}
	err := r.Database.GetContext(ctx, result, QueryGetExampleData, id)
	if err != nil {
		slog.ErrorContext(ctx, errs.ErrTagGetExampleDataDB,
			err,
			slog.Any("id", id),
		)
		return nil, errs.GenerateError(http.StatusInternalServerError, err)
	}

	return result, nil
}
