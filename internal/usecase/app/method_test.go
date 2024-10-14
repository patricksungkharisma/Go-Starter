package app_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/patricksungkharisma/go-starter/internal/config"
	"github.com/patricksungkharisma/go-starter/internal/entity"
	errs "github.com/patricksungkharisma/go-starter/internal/error"
	"github.com/patricksungkharisma/go-starter/internal/usecase/app"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_GetExampleData(t *testing.T) {
	var (
		ctrl        = gomock.NewController(t)
		ctx         = context.Background()
		mockAppRepo = NewMockAppRepo(ctrl)
		mockErr     = errors.New("testing error")
	)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req entity.GetExampleRequest
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(tt args)
		want    *entity.GetExampleResponse
		wantErr bool
	}{
		{
			name: "failed parsing ID",
			args: args{
				ctx: ctx,
				req: entity.GetExampleRequest{
					ID: "abc",
				},
			},
			mockFn:  func(tt args) {},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed get data from DB",
			args: args{
				ctx: ctx,
				req: entity.GetExampleRequest{
					ID: "1",
				},
			},
			mockFn: func(tt args) {
				mockAppRepo.EXPECT().GetExampleDataDB(gomock.Any(), gomock.Eq(int64(1))).Return(nil, errs.GenerateError(http.StatusInternalServerError, mockErr))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: entity.GetExampleRequest{
					ID: "1",
				},
			},
			mockFn: func(tt args) {
				mockAppRepo.EXPECT().GetExampleDataDB(gomock.Any(), gomock.Eq(int64(1))).Return(&entity.GetExampleResponse{
					ID:       1,
					Username: "Tester",
					IsActive: true,
				}, nil)
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
			u := &app.Usecase{
				Config:  cfg,
				AppRepo: mockAppRepo,
			}

			got, err := u.GetExampleData(tt.args.ctx, tt.args.req)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}
