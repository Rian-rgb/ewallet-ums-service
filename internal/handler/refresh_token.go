package handler

import (
	"ewallet-ums/internal/domain/auth"
	"ewallet-ums/internal/dto/auth_dto"
	"ewallet-ums/internal/errors"
	appErrors "github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/Rian-rgb/ewallet-common-lib/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshTokenHandler struct {
	RefreshTokenSvc auth.IRefreshTokenService
}

// @Summary     Refresh token user
// @Description Generates a new access token using a valid refresh token.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Bearer <token>"
// @Success      200      {object}  response.SuccessResponse{data=auth_dto.RefreshTokenResponse}
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /auth/refresh-token [put]
func (api *RefreshTokenHandler) RefreshToken(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	codeInternalError := appErrors.ErrCodeInternalServerError

	userData, ok := security.GetGinToken(ctx)
	if !ok {
		logger.WithContext(ctx).Error("token userData not found in context")
		response.SendError(
			ctx,
			codeInternalError.ToHTTPStatus(),
			codeInternalError,
			response.InternalServerErrorMessage,
		)
		return
	}

	result, err := api.RefreshTokenSvc.RefreshToken(ctx, token, userData)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := auth_dto.RefreshTokenResponse{Token: result}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, resp)
}
