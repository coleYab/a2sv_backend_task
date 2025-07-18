package router

import (
	"task_manager/controllers"
	"task_manager/data"

	"github.com/gin-gonic/gin"
)

type Router struct {
	c controllers.ITaskController
	e *gin.Engine
}

func New() *Router {
	taskService := data.New()
	taskController := controllers.New(taskService)
	return &Router{
		c: taskController,
		e: gin.Default(),
	}
}

func (r *Router)Run(addr string) {
	r.e.Run(addr)
}

func (r *Router)RegisterRoutes() {
	r.e.GET("/task", r.c.GetTasks)
	r.e.GET("/task/:id", r.c.GetTaskById)
	r.e.PUT("/task/:id", r.c.UpdateTask)
	r.e.DELETE("/task/:id", r.c.DeleteTask)
	r.e.POST("/task", r.c.CreateTask)
}