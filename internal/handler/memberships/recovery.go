package memberships

import (
	"net/http"

	"xprasetio/go-account-recovery.git/internal/constants"
	"xprasetio/go-account-recovery.git/internal/helpers"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RequestRecovery(c *gin.Context) {
	log := helpers.Logger

	var req memberships.ResetEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithError(err).Error("failed to parse request: ")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	if req.Email == "" {
		log.Error("Email is required")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, "Email is required", nil)
		return
	}

	err := h.service.InitiateRecovery(c.Request.Context(), req)
	if err != nil {
		log.WithError(err).Error("failed to initiate recovery: ")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusAccepted, constants.SuccessMessage, nil)
}
