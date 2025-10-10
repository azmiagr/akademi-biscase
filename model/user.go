package model

import "github.com/google/uuid"

type UserParam struct {
	UserID   uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
	RoleID   int       `json:"-"`
}

type UserRegister struct {
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type VerifyUser struct {
	UserID  uuid.UUID `json:"user_id" binding:"required"`
	OtpCode string    `json:"otp_code" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type GetUserProfileResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
