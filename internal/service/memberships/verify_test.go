package memberships

import (
	"context"
	"errors"
	"testing"

	"xprasetio/go-account-recovery.git/internal/configs"

	"go.uber.org/mock/gomock"
)

func Test_service_VerifyRecoveryCode(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)
	ctx := context.Background()
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mockFn  func(args args)
	}{
		// {
		// 	name: "success",
		// 	args: args{
		// 		ctx:  ctx,
		// 		code: "123456",
		// 	},
		// 	want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjI3MjM4MTksImlkIjo1LCJ1c2VybmFtZSI6ImFkbWluIn0.7XZJ0X3l0xWVh3kHkOq7JkIq2kW0XZLhW3kOq7JkIq2kW0XZLhW3kOq7J",
		// 	wantErr: false,
		// 	mockFn: func(args args) {
		// 		mockRepo.EXPECT().FindByRecoveryCode(args.ctx, args.code).Return(&memberships.User{
		// 			Model: gorm.Model{
		// 				ID: 5,
		// 			},
		// 			Username: "admin",
		// 		}, nil)
		// 	},
		// },
		// {
		// 	name: "invalid token",
		// 	args: args{
		// 		ctx:  ctx,
		// 		code: "invalid_token",
		// 	},
		// 	want:    "",
		// 	wantErr: true,
		// 	mockFn: func(args args) {
		// 		mockRepo.EXPECT().FindByRecoveryCode(args.ctx, args.code).Return(nil, errors.New("record not found"))
		// 	},
		// },
		{
			name: "error_find_by_recovery_code",
			args: args{
				ctx:  ctx,
				code: "123456",
			},
			want:    "",
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().FindByRecoveryCode(args.ctx, args.code).Return(nil, errors.New("database error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			got, err := s.VerifyRecoveryCode(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.VerifyRecoveryCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.VerifyRecoveryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

