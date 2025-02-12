package memberships

import (
	"context"
	"xprasetio/go-account-recovery.git/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) ResetPassword(ctx context.Context,token, newPassword string) error {
    userID, _, err := jwt.ValidateToken(token, s.cfg.Service.SecretKey)
    if err != nil {
        return err
    }
    user, err := s.repository.FindByID(ctx,userID)
    if err != nil {
        return err
    }
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user.Password = string(hashedPassword)
    user.RecoverCode = "coderecover"

    return s.repository.UpdateUser(ctx,user)
}
