package app

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	gomock "github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/patricksungkharisma/go-starter/internal/config"
)

func TestNew(t *testing.T) {
	var (
		ctrl         = gomock.NewController(t)
		mockDB, _, _ = sqlmock.New()
		db           = sqlx.NewDb(mockDB, "sqlmock")
	)
	defer ctrl.Finish()

	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name string
		args args
		want *Repo
	}{
		{
			name: "success",
			args: args{
				cfg: config.Config{},
			},
			want: &Repo{
				cfg:      config.Config{},
				database: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg, db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
