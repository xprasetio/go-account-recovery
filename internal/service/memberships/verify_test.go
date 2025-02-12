package memberships

import (
	"context"
	"testing"

	"xprasetio/go-account-recovery.git/internal/configs"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"

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
		{
			name: "error_find_by_recovery_code",
			args: args{
				ctx:  ctx,
				code: "123456",
			},
			want:    "",
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					GetUser(
						gomock.Eq(args.ctx),  // Context should match
						gomock.Any(),         // Empty string
						gomock.Any(),         // Empty string
						gomock.Any(),         // Any integer
						gomock.Eq(args.code), // Recovery code should match
					).
					Return(nil, gorm.ErrRecordNotFound)
			},
		},
		{
			name: "error_user_not_found",
			args: args{
				ctx:  ctx,
				code: "invalid_code",
			},
			want:    "",
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					GetUser(
						gomock.Eq(args.ctx),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Eq(args.code),
					).
					Return(nil, gorm.ErrRecordNotFound)
			},
		},
		{
			name: "success_verify_recovery_code",
			args: args{
				ctx:  ctx,
				code: "coderecover",
			},
			want:    "",  // We don't need to check exact token value
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().
					GetUser(
						gomock.Eq(args.ctx),
						gomock.Any(),
						gomock.Any(),
						gomock.Any(),
						gomock.Eq(args.code),
					).
					Return(&memberships.User{
						Model: gorm.Model{
							ID: 1,
						},
						Username:     "xprasetio",
						RecoverCode: "coderecover",
					}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{
					Service: configs.Service{
						SecretKey: "secret-key",
					},
				},
				repository: mockRepo,
			}
			got, err := s.VerifyRecoveryCode(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.VerifyRecoveryCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr { 
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}

