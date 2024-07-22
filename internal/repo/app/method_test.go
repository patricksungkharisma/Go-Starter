package app

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/patricksungkharisma/go-starter/internal/config"
	"github.com/patricksungkharisma/go-starter/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestRepo_GetExampleDataDB(t *testing.T) {
	var (
		ctrl               = gomock.NewController(t)
		ctx                = context.Background()
		cfg                = config.Config{}
		mockDB, mockSQL, _ = sqlmock.New()
		mockErr            = errors.New("testing error")
		db                 = sqlx.NewDb(mockDB, "sqlmock")
	)
	defer ctrl.Finish()

	type fields struct {
		cfg              config.Config
		databaseResource databaseResource
	}
	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFn  func(tt args)
		want    entity.GetExampleResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				cfg:              cfg,
				databaseResource: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			mockFn: func(tt args) {
				mockSQL.ExpectQuery(regexp.QuoteMeta(queryGetExampleData)).
					WithArgs(tt.id).
					WillReturnRows(sqlmock.
						NewRows([]string{"id", "username", "is_active"}).
						AddRow(1, "Tester", true))

			},
			want: entity.GetExampleResponse{
				ID:       1,
				Username: "Tester",
				IsActive: true,
			},
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				cfg:              cfg,
				databaseResource: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			mockFn: func(tt args) {
				mockSQL.ExpectQuery(regexp.QuoteMeta(queryGetExampleData)).
					WithArgs(tt.id).
					WillReturnError(mockErr)

			},
			want:    entity.GetExampleResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &Repo{
				cfg:      tt.fields.cfg,
				database: tt.fields.databaseResource,
			}

			got, err := r.GetExampleDataDB(ctx, tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.GetActivity() got error = %v, want error %v", err, tt.wantErr)
			}

			assert.Equal(t, got, tt.want)
		})
	}
}
