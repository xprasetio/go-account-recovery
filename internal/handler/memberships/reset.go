package memberships

import (
	"net/http"

	"xprasetio/go-account-recovery.git/internal/constants"
	"xprasetio/go-account-recovery.git/internal/helpers"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ResetPassword(c *gin.Context) {
    log := helpers.Logger
    bearerToken := c.GetHeader("Authorization")
    var req memberships.ResetPasswordRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.WithError(err).Error("failed to parse request: ")
        helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
        return
    }

    if err := h.service.ResetPassword(c.Request.Context(), bearerToken, req.Password); err != nil {
        log.WithError(err).Error("failed to reset password: ")
        helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
        return
    }

    helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}