package data

import (
	"fmt"
	"task_manager/controllers/dto"
	"task_manager/models"

	"github.com/google/uuid"
)

type ITaskService interface {
	GetTasks() []models.Task
	DeleteTask(id string) error
	GetTaskById(id string) (models.Task, error)
	CreateTask(task dto.CreateTaskDto) models.Task
	UpdateTask(id string, task dto.UpdateTaskDto) error
}

type TaskService struct {
	tasks map[string]models.Task
}

func New() *TaskService {
	return &TaskService{
		tasks: map[string]models.Task{},
	}
}

func (ts *TaskService) GetTasks() []models.Task {
	tasks := []models.Task{}

	for _, task := range ts.tasks {
		tasks = append(tasks, task)
	}	

	return tasks
}

func (ts *TaskService) GetTaskById(id string) (models.Task, error) {
	task, ok := ts.tasks[id]
	if !ok {
		return models.Task{}, fmt.Errorf("task not found")
	}

	return task, nil
}

func (ts *TaskService)CreateTask(task dto.CreateTaskDto) models.Task {
	newTask := models.Task{
		Id: uuid.NewString(),
		Title: task.Title,
		Description: task.Description,
		Status: task.Status,
		DueDate: task.DueDate,
	}

	ts.tasks[newTask.Id] = newTask

	return newTask
}

func (ts *TaskService)DeleteTask(id string) error {
	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf("task not found")
	} 

	delete(ts.tasks, id);
	return nil
}

func (ts *TaskService)UpdateTask(id string, task dto.UpdateTaskDto) error {
	olderTask, ok := ts.tasks[id]
	if !ok {
		return fmt.Errorf("task not found")
	}

	// update the task here
	olderTask.Title = task.Title
	olderTask.Description = task.Description
	olderTask.Status = task.Status
	olderTask.DueDate = task.DueDate

	ts.tasks[id] = olderTask

	return nil
}
