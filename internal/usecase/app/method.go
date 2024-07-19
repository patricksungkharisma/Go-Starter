package app

import (
	"context"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
)

func (u *Usecase) GetExampleData(ctx context.Context, req entity.GetExampleRequest) (entity.GetExampleResponse, error) {
	result, err := u.appRepo.GetExampleDataDB(ctx, req)
	if err != nil {
		return entity.GetExampleResponse{}, err
	}

	return result, nil
}
