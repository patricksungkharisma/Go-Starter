package app

import (
	"context"
	"strconv"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
)

func (u *Usecase) GetExampleData(ctx context.Context, req entity.GetExampleRequest) (entity.GetExampleResponse, error) {
	result := entity.GetExampleResponse{}

	parsedIDInt, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		return result, err
	}

	result, err = u.appRepo.GetExampleDataDB(ctx, parsedIDInt)
	if err != nil {
		return entity.GetExampleResponse{}, err
	}

	return result, nil
}
