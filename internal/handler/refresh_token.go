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
	RefreshTokenService auth.IRefreshTokenService
}

// Refresh godoc
// @Summary     Refresh token  user
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

	claim, ok := ctx.Get("token")
	if !ok {
		logger.WithContext(ctx).Error("token claim not found in context")
		response.SendError(
			ctx,
			codeInternalError.ToHTTPStatus(),
			codeInternalError,
			response.InternalServerErrorMessage,
		)
		return
	}

	tokenClaim, ok := claim.(*security.ClaimToken)
	if !ok {
		logger.WithContext(ctx).Error("failed to parse token claim")
		response.SendError(
			ctx,
			codeInternalError.ToHTTPStatus(),
			codeInternalError,
			response.InternalServerErrorMessage,
		)
		return
	}

	result, err := api.RefreshTokenService.RefreshToken(ctx.Request.Context(), token, *tokenClaim)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp := auth_dto.RefreshTokenResponse{Token: result}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, resp)
}
