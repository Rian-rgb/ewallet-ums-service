package errors

import (
	"errors"
	appErrors "github.com/Rian-rgb/ewallet-common-lib/errors"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/gin-gonic/gin"
)

func HandleServiceError(ctx *gin.Context, err error) {
	var (
		errCodeUnauthorized        = appErrors.ErrCodeUnauthorized
		errCodeInternalServerError = appErrors.ErrCodeInternalServerError
		errCodeUnknownError        = appErrors.ErrCodeUnknownError
	)
	if err == nil {
		return
	}

	switch {
	case errors.Is(err, ErrUserNotFound),
		errors.Is(err, ErrInvalidPassword),
		errors.Is(err, ErrInvalidToken):

		response.SendError(ctx, errCodeUnauthorized.ToHTTPStatus(), errCodeUnauthorized, err.Error())

	case errors.Is(err, ErrInternalServerError):
		response.SendError(ctx, errCodeInternalServerError.ToHTTPStatus(), errCodeInternalServerError, err.Error())

	default:
		response.SendError(ctx, errCodeUnknownError.ToHTTPStatus(), errCodeUnknownError, response.InternalServerErrorMessage)
	}
}
