package memberships

import (
	"context"
	"xprasetio/go-account-recovery.git/internal/models/memberships"
)

func (r *repository) CreateUser(ctx context.Context, model memberships.User) error {
	return r.db.Create(&model).Error
}

func (r *repository) GetUser(ctx context.Context, email, username string, id uint, recover_code string) (*memberships.User, error) {
	user := memberships.User{}
	res := r.db.Where("email = ?", email).Or("username = ?", username).Or("recover_code = ?", recover_code).Or("id = ?", id).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (r *repository) UpdateUser(ctx context.Context, user *memberships.User) error {
	return r.db.Save(user).Error
}

func (r *repository) FindByRecoveryCode(ctx context.Context, code string) (*memberships.User, error) {
    user  := memberships.User{}
    result := r.db.Where("recover_code = ?", code).First(&user)
    return &user, result.Error
}

func (r *repository) FindByID(ctx context.Context, id uint) (*memberships.User, error) {
    user :=memberships.User{}
    result := r.db.First(&user, id)
    return &user, result.Error
}