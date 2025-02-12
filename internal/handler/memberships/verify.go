package memberships

import (
	"net/http"

	"xprasetio/go-account-recovery.git/internal/constants"
	"xprasetio/go-account-recovery.git/internal/helpers"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"github.com/gin-gonic/gin"
)

func (h *Handler) VerifyRecoveryCode(c *gin.Context) {
    log := helpers.Logger
    var req memberships.ResetCodeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.WithError(err).Error("failed to parse request: ")
        helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
        return
    }

    token, err := h.service.VerifyRecoveryCode(c.Request.Context(), req.Code)
    if err != nil {
        log.WithError(err).Error("failed to verify recovery code: ")
        helpers.SendResponseHTTP(c, http.StatusUnauthorized, constants.ErrServerError, nil)
        return
    }

    helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, token)
}