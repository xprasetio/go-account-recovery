package memberships

import (
	"context"
	"xprasetio/go-account-recovery.git/internal/configs"
	"xprasetio/go-account-recovery.git/internal/models/memberships"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=memberships
type repository interface {
	CreateUser(ctx context.Context, model memberships.User) error
	GetUser(ctx context.Context,email, username string, id uint, recover_code string) (*memberships.User, error)
	UpdateUser(ctx context.Context,user *memberships.User) error
	FindByID(ctx context.Context,id uint) (*memberships.User, error) 
	FindByRecoveryCode(ctx context.Context,code string) (*memberships.User, error)

}

type service struct { 
	cfg *configs.Config // untuk generate key proses login
	repository repository
}

func NewService(cfg *configs.Config, repository repository) *service {
	return &service{
		cfg: cfg,
		repository: repository,
	}
}