package memberships

import (
	"context"
	"errors"
	"xprasetio/go-account-recovery.git/internal/models/memberships"
	"xprasetio/go-account-recovery.git/pkg/jwt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(ctx context.Context, request memberships.LoginRequest) (string, error) {
	userDetail, err := s.repository.GetUser(ctx,request.Email,"",0,"")
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Error get user from database")
		return "", err
	}
	if userDetail == nil {
		return "", errors.New("email not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password))
	if err != nil {		
		return "", errors.New("password or email not match")
	}

	accessToken, err := jwt.CreateToken(uint(userDetail.ID), userDetail.Username, s.cfg.Service.SecretKey)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate token")
		return "", err

	}
	return accessToken, nil
}