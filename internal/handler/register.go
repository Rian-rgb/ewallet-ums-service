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

type RegisterHandler struct {
	RegisterSvc auth.IRegisterService
}

// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account and securely stores the hashed password.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      auth_dto.RegisterRequest  true  "Register payload"
// @Success      200      {object}  response.SuccessResponse{data=auth_dto.RegisterResponse}
// @Failure      400      {object}  response.BadRequestResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /auth/register [post]
func (api *RegisterHandler) Register(ctx *gin.Context) {
	var (
		req            auth_dto.RegisterRequest
		resp           auth_dto.RegisterResponse
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

	userEntity := req.ToEntity()
	result, err := api.RegisterSvc.Register(ctx, userEntity)
	if err != nil {
		errors.HandleServiceError(ctx, err)
		return
	}

	resp = auth_dto.RegisterResponse{
		ID:          result.ID,
		Username:    result.Username,
		Email:       result.Email,
		PhoneNumber: result.PhoneNumber,
		Address:     result.Address,
		Dob:         result.Dob,
		FullName:    result.FullName,
	}

	response.SendSuccess(ctx, http.StatusCreated, response.SuccessMessage, resp)
}
