package app_test

import (
	"errors"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/patricksungkharisma/go-starter/internal/config"
	"github.com/patricksungkharisma/go-starter/internal/entity"
	errs "github.com/patricksungkharisma/go-starter/internal/error"
	"github.com/patricksungkharisma/go-starter/internal/handler/app"
	"github.com/patricksungkharisma/go-starter/internal/middleware"
	"github.com/patricksungkharisma/go-starter/internal/pkg"
)

func TestHandler_GetExample(t *testing.T) {
	var (
		ctrl           = gomock.NewController(t)
		mockAppUsecase = NewMockAppUsecase(ctrl)
		mockErr        = errors.New("testing error")
	)
	defer ctrl.Finish()

	type args struct{}
	tests := []struct {
		name   string
		args   args
		mockFn func(args args)
	}{
		{
			name: "failed",
			args: args{},
			mockFn: func(args args) {
				mockAppUsecase.EXPECT().GetExampleData(gomock.Any(), gomock.Any()).Return(nil, errs.GenerateError(http.StatusInternalServerError, mockErr))
			},
		},
		{
			name: "success",
			args: args{},
			mockFn: func(args args) {
				mockAppUsecase.EXPECT().GetExampleData(gomock.Any(), gomock.Any()).Return(&entity.GetExampleResponse{
					ID:       1,
					Username: "test",
					IsActive: true,
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			h := &app.Handler{
				Config:     config.Config{},
				AppUsecase: mockAppUsecase,
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
