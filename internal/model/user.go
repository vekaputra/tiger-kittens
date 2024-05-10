package model

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,max=64,email"`
	Password string `json:"password" validate:"required,min=8,max=20,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789"`
	Username string `json:"username" validate:"required,min=6,max=64,alphanum"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,min=6,max=64"`
	Password string `json:"password" validate:"required,min=8,max=20,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	Timestamp   string `json:"timestamp"`
}
