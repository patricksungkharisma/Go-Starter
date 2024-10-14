package app

import (
	"context"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
	errs "github.com/patricksungkharisma/go-starter/internal/error"
)

//go:generate mockgen -source=dependency.go -package=app_test -destination=dependency_mock_test.go
type AppRepo interface {
	GetExampleDataDB(ctx context.Context, id int64) (*entity.GetExampleResponse, *errs.TemplateError)
}
