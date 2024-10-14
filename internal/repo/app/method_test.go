package app_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	gomock "github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/patricksungkharisma/go-starter/internal/config"
	entity "github.com/patricksungkharisma/go-starter/internal/entity"
	"github.com/patricksungkharisma/go-starter/internal/repo/app"
	"github.com/stretchr/testify/assert"
)

func TestRepo_GetExampleDataDB(t *testing.T) {
	var (
		ctrl               = gomock.NewController(t)
		ctx                = context.Background()
		mockDB, mockSQL, _ = sqlmock.New()
		db                 = sqlx.NewDb(mockDB, "sqlmock")
		mockErr            = errors.New("testing error")
	)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		args    args
		mockFn  func(tt args)
		want    *entity.GetExampleResponse
		wantErr bool
	}{
		{
			name: "failed",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mockFn: func(tt args) {
				mockSQL.ExpectQuery(regexp.QuoteMeta(app.QueryGetExampleData)).
					WithArgs(tt.id).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			mockFn: func(tt args) {
				mockSQL.ExpectQuery(regexp.QuoteMeta(app.QueryGetExampleData)).
					WithArgs(tt.id).
					WillReturnRows(sqlmock.
						NewRows([]string{"id", "username", "is_active"}).
						AddRow(1, "Tester", true))

			},
			want: &entity.GetExampleResponse{
				ID:       1,
				Username: "Tester",
				IsActive: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			cfg := config.Config{}
			r := &app.Repo{
				Config:   cfg,
				Database: db,
			}

			got, err := r.GetExampleDataDB(ctx, tt.args.id)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}
