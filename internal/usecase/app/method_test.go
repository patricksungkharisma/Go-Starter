package app

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/patricksungkharisma/go-starter/internal/config"
	"github.com/patricksungkharisma/go-starter/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_GetExampleData(t *testing.T) {
	var (
		ctrl        = gomock.NewController(t)
		ctx         = context.Background()
		cfg         = config.Config{}
		mockAppRepo = NewMockappRepo(ctrl)
		mockErr     = errors.New("testing error")
	)
	defer ctrl.Finish()

	type fields struct {
		cfg     config.Config
		appRepo appRepo
	}
	type args struct {
		ctx context.Context
		req entity.GetExampleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFn  func()
		want    entity.GetExampleResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				cfg:     cfg,
				appRepo: mockAppRepo,
			},
			args: args{
				ctx: ctx,
				req: entity.GetExampleRequest{
					ID: "1",
				},
			},
			mockFn: func() {
				mockAppRepo.EXPECT().GetExampleDataDB(gomock.Any(), gomock.Eq(int64(1))).Return(entity.GetExampleResponse{
					ID:       1,
					Username: "Tester",
					IsActive: true,
				}, nil)
			},
			want: entity.GetExampleResponse{
				ID:       1,
				Username: "Tester",
				IsActive: true,
			},
			wantErr: false,
		},
		{
			name: "failed parsing ID",
			fields: fields{
				cfg:     cfg,
				appRepo: mockAppRepo,
			},
			args: args{
				ctx: ctx,
				req: entity.GetExampleRequest{
					ID: "abc",
				},
			},
			mockFn:  func() {},
			want:    entity.GetExampleResponse{},
			wantErr: true,
		},
		{
			name: "failed get data from DB",
			fields: fields{
				cfg:     cfg,
				appRepo: mockAppRepo,
			},
			args: args{
				ctx: ctx,
				req: entity.GetExampleRequest{
					ID: "1",
				},
			},
			mockFn: func() {
				mockAppRepo.EXPECT().GetExampleDataDB(gomock.Any(), gomock.Eq(int64(1))).Return(entity.GetExampleResponse{}, mockErr)
			},
			want:    entity.GetExampleResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			u := &Usecase{
				cfg:     tt.fields.cfg,
				appRepo: tt.fields.appRepo,
			}
			got, err := u.GetExampleData(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetExampleData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equalf(t, tt.want, got, "GetExampleData() got = %v, want %v", got, tt.want)
		})
	}
}
