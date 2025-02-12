package memberships

import (
	"context"
	"errors"
	"testing"

	"xprasetio/go-account-recovery.git/internal/configs"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"gorm.io/gorm"

	"go.uber.org/mock/gomock"
)

func Test_service_InitiateRecovery(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)
	ctx := context.Background()
	type args struct {
		ctx     context.Context
		request memberships.ResetEmailRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "error_email_not_found",
			args: args{
				ctx: ctx,
				request: memberships.ResetEmailRequest{
					Email: "xprasetio@gmail.com",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				// Mock GetUser returns error
				mockRepo.EXPECT().
					GetUser(args.ctx, args.request.Email, "", uint(0), "").
					Return(nil, gorm.ErrRecordNotFound)
			},
		},		
		{
			name: "error_get_user_from_database",
			args: args{
				ctx: ctx,
				request: memberships.ResetEmailRequest{
					Email: "xprasetio@gmail.com",
				},
			},
			wantErr: true,	
			mockFn: func(args args) {
				// Mock GetUser returns database error
				mockRepo.EXPECT().
					GetUser(args.ctx, args.request.Email, "", uint(0), "").
					Return(nil, errors.New("database error"))
			},			
		},
		{
			name: "error_update_user",
			args: args{
				ctx: ctx,
				request: memberships.ResetEmailRequest{
					Email: "xprasetio@gmail.com",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				user := &memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "xprasetio@gmail.com",
					Username: "xprasetio",
				}
				// Mock GetUser success
				mockRepo.EXPECT().
					GetUser(args.ctx, args.request.Email, "", uint(0), "").
					Return(user, nil)

				// Mock UpdateUser fails
				mockRepo.EXPECT().
					UpdateUser(args.ctx, gomock.Any()).
					Return(errors.New("failed to update user"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:      &configs.Config{},
				repository: mockRepo,
			}
			if err := s.InitiateRecovery(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.InitiateRecovery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
