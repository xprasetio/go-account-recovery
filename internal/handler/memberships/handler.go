package memberships

import (
	"context"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"github.com/gin-gonic/gin"
)


type service interface {
	SignUp(ctx context.Context, request memberships.SignUpRequest) error
	Login(ctx context.Context, request memberships.LoginRequest) (string, error)
	InitiateRecovery(ctx context.Context, request memberships.ResetEmailRequest) error
	VerifyRecoveryCode(ctx context.Context,code string) (string, error)
	ResetPassword(ctx context.Context,token, newPassword string) error
}

type Handler struct {
	 *gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/api/v1/memberships")
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)
	route.POST("/recovery", h.RequestRecovery)
	route.POST("/verify", h.VerifyRecoveryCode)
	route.POST("/reset-password", h.ResetPassword)
}