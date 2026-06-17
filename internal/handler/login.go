package handler

import (
	"ewallet-ums/internal/domain/auth"
	"ewallet-ums/internal/dto/auth_dto"
	"ewallet-ums/internal/errors"
	appErrors "github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	LoginService auth.ILoginService
}

// Login godoc
// @Summary      Login user
// @Description  Authenticates user credentials (username and password).
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      auth_dto.LoginRequest  true  "Payload login"
// @Success      200      {object}  response.SuccessResponse{data=auth_dto.LoginResponse}
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /auth/login [post]
func (api *LoginHandler) Login(ctx *gin.Context) {
	var (
		req            auth_dto.LoginRequest
		resp           auth_dto.LoginResponse
		codeBadRequest = appErrors.ErrCodeBadRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.WithContext(ctx).Error("failed to parse JSON request: ", err)
		response.SendBadRequest(ctx, codeBadRequest, response.InvalidJSONFormatMessage, nil)
		return
	}

	errFields := req.Validate()
	if errFields != nil {
		logger.WithContext(ctx).Warn("request body validation failed")
		response.SendBadRequest(ctx, codeBadRequest, response.InvalidRequestMessage, errFields)
		return
	}

	result, token, refreshToken, err := api.LoginService.Login(ctx, req.Username, req.Password)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp = auth_dto.LoginResponse{
		UserID:       result.ID,
		Username:     result.Username,
		FullName:     result.FullName,
		Email:        result.FullName,
		Token:        token,
		RefreshToken: refreshToken,
	}

	response.SendSuccess(ctx, http.StatusOK, response.SuccessMessage, resp)
}
