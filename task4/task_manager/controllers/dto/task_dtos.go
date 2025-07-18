package dto

import "time"

type CreateTaskDto struct {
	Title       string    `json:"title" binding:"required,min=3,max=100"`
	Description string    `json:"description" binding:"required,min=5,max=500"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
	Status      string    `json:"status" binding:"required,oneof=pending in-progress completed"`
}

type UpdateTaskDto struct {
	Title       string    `json:"title" binding:"required,min=3,max=100"`
	Description string    `json:"description" binding:"required,min=5,max=500"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
	Status      string    `json:"status" binding:"required,oneof=pending in-progress completed"`
}
