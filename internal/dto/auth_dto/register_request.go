package auth_dto

import (
	"ewallet-ums/internal/domain/user"
	"github.com/Rian-rgb/ewallet-common-lib/response"
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Username    string `json:"username" validate:"required,max=20"`
	Email       string `json:"email" validate:"required,email,max=100"`
	PhoneNumber string `json:"phone_number" validate:"required,max=15"`
	Address     string `json:"address"`
	Dob         string `json:"dob"`
	Password    string `json:"password" validate:"required,min=8,max=255"`
	FullName    string `json:"full_name" validate:"required,max=100"`
}

func (req RegisterRequest) Validate() []response.ValidationErrorField {
	v := validator.New()
	err := v.Struct(req)

	return response.MapValidationErrors(err)
}

func (req *RegisterRequest) ToEntity() *user.Entity {
	return &user.Entity{
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		Address:     req.Address,
		Dob:         req.Dob,
		FullName:    req.FullName,
	}
}
