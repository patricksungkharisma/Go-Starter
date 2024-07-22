package app

import (
	"context"

	entity "github.com/patricksungkharisma/go-starter/internal/entity"
)

//go:generate mockgen -source=dependency.go -package=app -destination=dependency_mock_test.go
type appRepo interface {
	GetExampleDataDB(ctx context.Context, id int64) (entity.GetExampleResponse, error)
}
