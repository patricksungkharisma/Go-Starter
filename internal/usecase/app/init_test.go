package app

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/patricksungkharisma/go-starter/internal/config"
)

func TestNew(t *testing.T) {
	var (
		ctrl        = gomock.NewController(t)
		mockAppRepo = NewMockappRepo(ctrl)
	)
	defer ctrl.Finish()

	type args struct {
		cfg     config.Config
		appRepo appRepo
	}
	tests := []struct {
		name string
		args args
		want *Usecase
	}{
		{
			name: "success",
			args: args{
				cfg:     config.Config{},
				appRepo: mockAppRepo,
			},
			want: &Usecase{
				cfg:     config.Config{},
				appRepo: mockAppRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg, tt.args.appRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
