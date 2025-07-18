package controllers

import (
	"net/http"
	"task_manager/controllers/dto"
	"task_manager/data"

	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	CreateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	GetTasks(c *gin.Context)
	GetTaskById(c *gin.Context)
}

type TaskController struct {
	taskService data.ITaskService
}

func New(service data.ITaskService) *TaskController {
	return &TaskController{
		taskService : service,
	}
}

func (c *TaskController)CreateTask(ctx *gin.Context) {
	var createTaskDto dto.CreateTaskDto

	if err := ctx.ShouldBind(&createTaskDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	task := c.taskService.CreateTask(createTaskDto)

	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskController)DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	if err := c.taskService.DeleteTask(taskId); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err.Error(),
		});
		return 
	}
	ctx.Status(http.StatusNoContent)
}

func (c *TaskController)UpdateTask(ctx *gin.Context) {
	var updateTaskDto dto.UpdateTaskDto
	taskId := ctx.Param("id")

	if err := ctx.ShouldBind(&updateTaskDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return 
	}

	if err := c.taskService.UpdateTask(taskId, updateTaskDto); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, _ := c.taskService.GetTaskById(taskId)

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController)GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.taskService.GetTasks())
}

func (c *TaskController)GetTaskById(ctx *gin.Context) {
	taskId := ctx.Param("id")
	task, err := c.taskService.GetTaskById(taskId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, task)
}