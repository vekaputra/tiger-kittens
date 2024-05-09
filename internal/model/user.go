package model

type RegisterUserRequest struct {
	Email    string `json:"email"`    // validate:"required,email"`
	Password string `json:"password"` // validate:"min=8,max=20,regexp=^(?=.*[A-Z])(?=.*[!@#$&*])(?=.*[0-9])(?=.*[a-z])$"`
	Username string `json:"username"` // validate:"required,min=6,max=30,alphanum"`
}
