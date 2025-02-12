package memberships

import (
	"net/http"
	"xprasetio/go-account-recovery.git/internal/constants"
	"xprasetio/go-account-recovery.git/internal/helpers"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"github.com/gin-gonic/gin"
)


func (h *Handler) Login(c *gin.Context) {
	log := helpers.Logger

	var req memberships.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithError(err).Error("failed to parse request: ")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	accessToken, err := h.service.Login(c.Request.Context(), req)

	if err != nil { 
		log.WithError(err).Error("failed to login: ")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest,nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, accessToken)
}
