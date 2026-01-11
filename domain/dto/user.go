package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	UUID       uuid.UUID `json:"uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role       string `json:"role"`
	PhoneNumber string `json:"phone_number"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type RegisterRequest struct {
	Name      string `json:"name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Email     string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Role     uint
}	


type RegisterResponse struct {
	User  UserResponse `json:"user"`
}

type UpdateRequest struct {
	Name      string `json:"name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
	Email     string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Role     uint
}

