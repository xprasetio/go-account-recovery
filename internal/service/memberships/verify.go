package memberships

import (
	"context"
	"xprasetio/go-account-recovery.git/pkg/jwt"

	"github.com/rs/zerolog/log"
)

func (s *service) VerifyRecoveryCode(ctx context.Context,code string) (string, error) {
    userDetail, err := s.repository.GetUser(ctx,"", "", 0, code)
    if err != nil {
		log.Error().Err(err).Msg("Error get user from database")
        return "", err
    }
    // Generate JWT token
   accessToken, err := jwt.CreateToken(uint(userDetail.ID), userDetail.Username, s.cfg.Service.SecretKey)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate token")
		return "", err

	}
	return accessToken, nil
}