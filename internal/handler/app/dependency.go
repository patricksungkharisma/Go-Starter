package app

import (
	"context"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
	errs "github.com/patricksungkharisma/go-starter/internal/error"
)

//go:generate mockgen -source=dependency.go -package=app_test -destination=dependency_mock_test.go
type AppUsecase interface {
	GetExampleData(ctx context.Context, req entity.GetExampleRequest) (*entity.GetExampleResponse, *errs.TemplateError)
}
