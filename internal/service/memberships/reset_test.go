package memberships

import (
	"context"
	"errors"
	"testing"

	"xprasetio/go-account-recovery.git/internal/configs"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"go.uber.org/mock/gomock"
)

func Test_service_ResetPassword(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)
	ctx := context.Background()
	type args struct {
		ctx         context.Context
		token       string
		newPassword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		// {
		// 	name: "success_reset_password",
		// 	args: args{
		// 		ctx:         ctx,
		// 		token:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkzNDQxMDksImlkIjoxNSwidXNlcm5hbWUiOiJ4cHJhc2V0aW8ifQ.MCEf3jd_xkOKGeL3s6AFgErl6ErQd3zMcCQG-px0C8o",
		// 		newPassword: "newpassword123",
		// 	},
		// 	wantErr: false,
		// 	mockFn: func(args args) {
		// 		userID := uint(1)   			
		// 		// Mock FindByID
		// 		mockRepo.EXPECT().
		// 			FindByID(ctx, userID).
		// 			Return(&memberships.User{
		// 				Model: gorm.Model{ID: userID},
		// 				Email: "xprasetio@gmail.com",
		// 				Password: "kocak789",
		// 			}, nil)

		// 		// Mock UpdateUser dengan password yang sudah di-hash
		// 		mockRepo.EXPECT().
		// 			UpdateUser(args.ctx,gomock.Any()).
		// 			Return(nil)
		// 	},
   		// },
		{	
			name:    "Invalid token",
			args:    args{ctx, "invalid_token", "new_password"},
			wantErr: true,
			mockFn: func(args args) {},
		},
		{
			name:    "User not found",
			args:    args{ctx, "valid_token", "new_password"},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().FindByID(args.ctx, gomock.Any()).Return(nil, errors.New("user not found"))
			},
		},

		{
			name:    "Failed to hash password",
			args:    args{ctx, "valid_token", ""},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().FindByID(args.ctx, gomock.Any()).Return(&memberships.User{}, nil)
			},
		},
		{
			name:    "Failed to update user",
			args:    args{ctx, "valid_token", "new_password"},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().FindByID(args.ctx, gomock.Any()).Return(&memberships.User{}, nil)
				mockRepo.EXPECT().UpdateUser(args.ctx, gomock.Any()).Return(errors.New("update failed"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cfg:        &configs.Config{
					Service: configs.Service{
						SecretKey: "secret_key",
					},
				},
				repository: mockRepo,
			}
			if err := s.ResetPassword(tt.args.ctx, tt.args.token, tt.args.newPassword); (err != nil) != tt.wantErr {
				t.Errorf("service.ResetPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
