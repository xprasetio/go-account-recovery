package memberships

import (
	"net/http"
	"xprasetio/go-account-recovery.git/internal/constants"
	"xprasetio/go-account-recovery.git/internal/helpers"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var (
		log = helpers.Logger
	)
	var req memberships.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(ctx, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	req.RecoverCode = "coderecover"
	err := h.service.SignUp(ctx.Request.Context(), req)
	if err != nil {
		log.Error("failed to sign up: ", err)
		helpers.SendResponseHTTP(ctx, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponseHTTP(ctx, http.StatusCreated, constants.SuccessMessage, nil)
}