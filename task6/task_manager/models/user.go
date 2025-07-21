package models

type User struct {
	Id       string
	UserName string
	Password string
	Role     string
}

const (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)
