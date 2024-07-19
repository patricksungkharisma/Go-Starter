package app

import (
	"context"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
)

func (r *Repo) GetExampleDataDB(ctx context.Context, req entity.GetExampleRequest) (entity.GetExampleResponse, error) {
	result := entity.GetExampleResponse{}

	err := r.database.GetContext(ctx, &result, queryGetExampleData, req.ID)
	if err != nil {
		return result, err
	}

	return result, nil
}
