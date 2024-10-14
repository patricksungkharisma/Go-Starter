package app_test

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/patricksungkharisma/go-starter/internal/config"
	"github.com/patricksungkharisma/go-starter/internal/usecase/app"
)

func TestNew(t *testing.T) {
	var (
		ctrl        = gomock.NewController(t)
		mockAppRepo = NewMockAppRepo(ctrl)
	)
	defer ctrl.Finish()

	type args struct {
		cfg     config.Config
		appRepo app.AppRepo
	}
	tests := []struct {
		name string
		args args
		want *app.Usecase
	}{
		{
			name: "success",
			args: args{
				cfg:     config.Config{},
				appRepo: mockAppRepo,
			},
			want: &app.Usecase{
				Config:  config.Config{},
				AppRepo: mockAppRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := app.New(tt.args.cfg, tt.args.appRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
