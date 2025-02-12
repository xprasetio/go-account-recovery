package memberships

import (
	"context"
	"errors"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"gorm.io/gorm"

	"github.com/rs/zerolog/log"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, request memberships.SignUpRequest) error {
	existingUser, err := s.repository.GetUser(ctx, request.Email, request.Username, uint(0),"")

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("failed to get user")
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
		return err
	}

	model := memberships.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  string(pass),
		RecoverCode: request.RecoverCode,
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}
	err = s.repository.CreateUser(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		return err
	}
	return nil
}