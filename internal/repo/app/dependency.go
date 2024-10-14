package app

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=dependency.go -package=app_test -destination=dependency_mock_test.go
type DatabaseResource interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}
