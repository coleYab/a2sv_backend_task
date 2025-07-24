package controllers

import (
	"fmt"
	"net/http"
	"task_manager/delivery/controllers/dto"
	"task_manager/domain"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	CreateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	GetTasks(c *gin.Context)
	GetTaskByID(c *gin.Context)
}

type TaskController struct {
	taskUseCase usecases.ITaskUseCase
}

func NewTaskController(uc usecases.ITaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: uc,
	}
}

func verifyAdmin(ctx *gin.Context) error {
	user, err := getUserFromContext(ctx)
	if err != nil {
		return err
	}

	if user.Role == domain.UserRoleAdmin {
		return nil
	}

	return fmt.Errorf("forbidden resource")
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var createTaskDto dto.CreateTaskDto

	if err := verifyAdmin(ctx); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBind(&createTaskDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task := c.taskUseCase.CreateTask(createTaskDto)

	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	if err := verifyAdmin(ctx); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	taskID := ctx.Param("id")
	if err := c.taskUseCase.DeleteTask(taskID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	if err := verifyAdmin(ctx); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updateTaskDto dto.UpdateTaskDto
	taskID := ctx.Param("id")
	if err := ctx.ShouldBind(&updateTaskDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.taskUseCase.UpdateTask(taskID, updateTaskDto); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, _ := c.taskUseCase.GetTaskByID(taskID)

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.taskUseCase.GetTasks())
}

func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	taskID := ctx.Param("id")
	task, err := c.taskUseCase.GetTaskByID(taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, task)
}
