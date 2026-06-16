package handler

import (
	"ewallet-ums/internal/domain/auth"
	"ewallet-ums/internal/errors"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogoutHandler struct {
	LogoutService auth.ILogoutService
}

// Logout godoc
// @Summary      Logout user
// @Description  Logs out the authenticated user and invalidates the refresh token
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Bearer <token>"
// @Success      200      {object}  response.SuccessResponse
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /user/logout [delete]
func (api *LogoutHandler) Logout(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	err := api.LogoutService.Logout(ctx, token)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, nil)
}
