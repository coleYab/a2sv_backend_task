// Package domain: domain module
package domain

import "time"

type User struct {
	ID       string
	UserName string
	Password string
	Role     string
}

const (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Status      string    `json:"status"`
}
