package app

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/patricksungkharisma/go-starter/internal/config"
	"github.com/patricksungkharisma/go-starter/internal/entity"
	"github.com/patricksungkharisma/go-starter/internal/middleware"
	"github.com/patricksungkharisma/go-starter/internal/pkg"
)

func TestHandler_GetExample(t *testing.T) {
	var (
		ctrl           = gomock.NewController(t)
		cfg            = config.Config{}
		mockAppUsecase = NewMockappUsecase(ctrl)
		mockErr        = errors.New("testing error")
	)
	defer ctrl.Finish()

	type fields struct {
		cfg        config.Config
		appUsecase appUsecase
	}
	type args struct{}
	tests := []struct {
		name   string
		fields fields
		args   args
		mockFn func(args args)
	}{
		{
			name: "success",
			fields: fields{
				cfg:        cfg,
				appUsecase: mockAppUsecase,
			},
			args: args{},
			mockFn: func(args args) {
				mockAppUsecase.EXPECT().GetExampleData(gomock.Any(), gomock.Any()).Return(entity.GetExampleResponse{
					ID:       1,
					Username: "Test",
					IsActive: true,
				}, nil)
			},
		},
		{
			name: "failed",
			fields: fields{
				cfg:        cfg,
				appUsecase: mockAppUsecase,
			},
			args: args{},
			mockFn: func(args args) {
				mockAppUsecase.EXPECT().GetExampleData(gomock.Any(), gomock.Any()).Return(entity.GetExampleResponse{}, mockErr)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			h := &Handler{
				cfg:        tt.fields.cfg,
				appUsecase: tt.fields.appUsecase,
			}

			uriParam := map[string]string{}
			c := pkg.CreateContextMock("GET", "localhost", nil, &uriParam, nil, nil)

			r := &middleware.RequestContext{
				GinContext: c,
			}

			h.GetExample(r)
		})
	}
}
