package model

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Name   string `json:"name"`
	UserId int    `json:"userId"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}
