// Package usecases: usecases module
package usecases

import (
	"task_manager/delivery/controllers/dto"
	"task_manager/domain"
	"task_manager/repositories"
)

type ITaskUseCase interface {
	GetTasks() []domain.Task
	DeleteTask(id string) error
	GetTaskByID(id string) (domain.Task, error)
	CreateTask(task dto.CreateTaskDto) domain.Task
	UpdateTask(id string, task dto.UpdateTaskDto) error
}

type TaskUseCase struct {
	taskRepository repositories.ITaskRepository
}

func NewTaskUseCase(repo repositories.ITaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepository: repo,
	}
}

func (tu *TaskUseCase) GetTasks() []domain.Task {
	return tu.taskRepository.GetTasks()
}

func (tu *TaskUseCase) GetTaskByID(id string) (domain.Task, error) {
	return tu.taskRepository.GetTaskByID(id)
}

func (tu *TaskUseCase) CreateTask(task dto.CreateTaskDto) domain.Task {
	return tu.taskRepository.CreateTask(task)
}

func (tu *TaskUseCase) DeleteTask(id string) error {
	return tu.taskRepository.DeleteTask(id)
}

func (tu *TaskUseCase) UpdateTask(id string, task dto.UpdateTaskDto) error {
	return tu.taskRepository.UpdateTask(id, task)
}
