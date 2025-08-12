// Package dto
package dto

type RegisterUserDTO struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type UpdateProfileDTO struct {
	Bio            string `json:"bio" binding:"required"`
	ProfilePicture string `json:"profilePicture" binding:"required,url"`
}