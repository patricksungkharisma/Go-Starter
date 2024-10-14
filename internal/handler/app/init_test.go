package app_test

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/patricksungkharisma/go-starter/internal/config"
	"github.com/patricksungkharisma/go-starter/internal/handler/app"
)

func TestNew(t *testing.T) {
	var (
		ctrl           = gomock.NewController(t)
		mockAppUsecase = NewMockAppUsecase(ctrl)
	)
	defer ctrl.Finish()

	type args struct {
		cfg        config.Config
		appUsecase app.AppUsecase
	}
	tests := []struct {
		name string
		args args
		want *app.Handler
	}{
		{
			name: "success",
			args: args{
				cfg:        config.Config{},
				appUsecase: mockAppUsecase,
			},
			want: &app.Handler{
				Config:     config.Config{},
				AppUsecase: mockAppUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := app.New(tt.args.cfg, tt.args.appUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
