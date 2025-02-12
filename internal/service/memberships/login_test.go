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

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		request memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				request: memberships.LoginRequest{
					Email:    "xprasetio@gmail.com",
					Password: "admin789",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(ctx,args.request.Email, "", uint(0),"").Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:     "xprasetio@gmail.com",
					Username:  "xprasetio",
					Password:  "$2a$10$/.q9g7OqWAAmym.lEq7mxe.3MkGOrxxTK.Zqh4vDi4u28PlbNENiO",
				},nil)
			},
		},
		{
			name: "failed when password not match",
			args: args{
				ctx: ctx,
				request: memberships.LoginRequest{
					Email:    "xprasetio@gmail.com",
					Password: "admin789",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(ctx,args.request.Email, "", uint(0),"").Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:     "xprasetio@gmail.com",
					Username:  "xprasetio",
					Password:  "wrong password bro",
				},nil)
			},
			
		},
		{
			name: "failed when get user",
			args: args{
				ctx: ctx,
				request: memberships.LoginRequest{
					Email:    "xprasetio@gmail.com",
					Password: "admin789",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(ctx,args.request.Email, "", uint(0),"").Return(nil, assert.AnError)
			},
			
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{
					Service: configs.Service{
						SecretKey: "secret",
					},
				},
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
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
