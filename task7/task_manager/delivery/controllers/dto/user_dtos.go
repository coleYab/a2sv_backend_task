package dto

type RegisterUserDto struct {
	UserName string `json:"userName" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

type LoginDto struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}
