package app

import (
	"context"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
)

//go:generate mockgen -source=dependency.go -package=app -destination=dependency_mock_test.go
type appUsecase interface {
	GetExampleData(ctx context.Context, req entity.GetExampleRequest) (entity.GetExampleResponse, error)
}
