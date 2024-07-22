package app

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/patricksungkharisma/go-starter/internal/config"
)

func TestNew(t *testing.T) {
	var (
		ctrl           = gomock.NewController(t)
		mockAppUsecase = NewMockappUsecase(ctrl)
	)
	defer ctrl.Finish()

	type args struct {
		cfg        config.Config
		appUsecase appUsecase
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "success",
			args: args{
				cfg:        config.Config{},
				appUsecase: mockAppUsecase,
			},
			want: &Handler{
				cfg:        config.Config{},
				appUsecase: mockAppUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg, tt.args.appUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
