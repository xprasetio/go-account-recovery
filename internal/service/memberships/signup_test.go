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

func Test_service_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)
	ctx := context.Background()
	type args struct {
		ctx     context.Context
		request memberships.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFunc func(args args)
	}{
		{
            name: "success",
            args: args{
                ctx: ctx,
                request: memberships.SignUpRequest{
                    Email:    "xprasetio@gmail.com",
                    Username: "xprasetio",
                    Password: "admin789",
                },
            },
            wantErr: false,
            mockFunc: func(args args) {
                mockRepo.EXPECT().GetUser(ctx,args.request.Email, args.request.Username, uint(0),"").Return(nil, gorm.ErrRecordNotFound).Times(1)
				mockRepo.EXPECT().CreateUser(ctx,gomock.Any()).Return(nil).Times(1)
            },
        },
		{
			name: "failed when get user",
			args: args{
				ctx: ctx,
				request: memberships.SignUpRequest{
					Email:     "xprasetio@gmail.com",
					Username:  "xprasetio",
					Password:  "admin789",
				},
			},
			wantErr: true,
			mockFunc: func(args args) {
				mockRepo.EXPECT().GetUser(ctx,args.request.Email, args.request.Username, uint(0),"").
				Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when create user",
			args: args{
				ctx: ctx,
				request: memberships.SignUpRequest{
					Email:     "xprasetio@gmail.com",
					Username:  "xprasetio",
					Password:  "admin789",
				},
			},
			wantErr: true,
			mockFunc: func(args args) {
				mockRepo.EXPECT().GetUser(ctx,args.request.Email, args.request.Username, uint(0),"").
				Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(ctx,gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.args)
			
			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			if err := s.SignUp(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
