package dto

type PromoteUserDTO struct {
	ID   string `json:"id" binding:"required"`
	Role string `json:"role" binding:"required,oneof=admin user"`
}
