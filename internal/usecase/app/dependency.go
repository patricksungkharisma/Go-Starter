package app

import (
	"context"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
)

type appRepo interface {
	GetExampleDataDB(ctx context.Context, req entity.GetExampleRequest) (entity.GetExampleResponse, error)
}
