package auth_dto

type RegisterResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Dob         string `json:"dob"`
	FullName    string `json:"full_name"`
}
