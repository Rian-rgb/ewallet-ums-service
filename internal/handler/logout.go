package handler

import (
	"ewallet-ums/internal/domain/auth"
	"ewallet-ums/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutHandler struct {
	LogoutSvc auth.ILogoutService
}

// @Summary		Logout User
// @Description	Logs out the authenticated user and invalidates the refresh token.
// @Tags		Auth
// @Accept		json
// @Produce		json
//
// @Param		Authorization	header	string	true	"Bearer <token>"
//
// @Success		200	{object}	response.SuccessResponse	"Success"
// @Failure		401	{object}	response.ErrorResponse		"Unauthorized"
// @Failure		500	{object}	response.ErrorResponse		"Internal Server Error"
//
// @Security	BearerAuth
// @Router		/auth/logout [delete]
func (hdl *LogoutHandler) Logout(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	err := hdl.LogoutSvc.Logout(ctx, token)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, nil)
}
