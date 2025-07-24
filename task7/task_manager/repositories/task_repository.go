package repositories

import (
	"context"
	"fmt"
	"task_manager/delivery/controllers/dto"
	"task_manager/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ITaskRepository interface {
	GetTasks() []domain.Task
	DeleteTask(id string) error
	GetTaskByID(id string) (domain.Task, error)
	CreateTask(task dto.CreateTaskDto) domain.Task
	UpdateTask(id string, task dto.UpdateTaskDto) error
}

type TaskRepository struct {
	collection *mongo.Collection
	tasks      map[string]domain.Task
}

func NewTaskRepository(collection *mongo.Collection) *TaskRepository {
	return &TaskRepository{
		collection: collection,
		tasks:      map[string]domain.Task{},
	}
}

func (ts *TaskRepository) GetTasks() []domain.Task {
	findOpts := options.Find()
	cursor, err := ts.collection.Find(context.TODO(), bson.D{{}}, findOpts)
	if err != nil {
		return nil
	}

	tasks := []domain.Task{}

	for cursor.Next(context.TODO()) {
		var task domain.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	cursor.Close(context.TODO())
	return tasks
}

func (ts *TaskRepository) GetTaskByID(id string) (domain.Task, error) {
	findOneOpts := options.FindOne()
	filter := bson.D{{Key: "id", Value: id}}

	var task domain.Task
	if err := ts.collection.FindOne(context.TODO(), filter, findOneOpts).Decode(&task); err != nil {
		return domain.Task{}, fmt.Errorf("find failed %v", err.Error())
	}

	return task, nil
}

func (ts *TaskRepository) CreateTask(task dto.CreateTaskDto) domain.Task {
	newTask := domain.Task{
		ID:          uuid.NewString(),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
	}

	ts.collection.InsertOne(context.TODO(), newTask)

	return newTask
}

func (ts *TaskRepository) DeleteTask(id string) error {
	findOneAndDeleteOpts := options.FindOneAndDelete()
	filters := bson.D{{Key: "id", Value: id}}
	var task domain.Task
	if err := ts.collection.FindOneAndDelete(context.TODO(), filters, findOneAndDeleteOpts).Decode(&task); err != nil {
		return err
	}

	return nil
}

func (ts *TaskRepository) UpdateTask(id string, task dto.UpdateTaskDto) error {
	oldtask, err := ts.GetTaskByID(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "id", Value: id}}
	findOneOpts := options.FindOneAndReplace()

	oldtask.Title = task.Title
	oldtask.Description = task.Description
	oldtask.DueDate = task.DueDate
	oldtask.Status = task.Status

	ts.collection.FindOneAndReplace(context.TODO(), filter, oldtask, findOneOpts)

	return nil
}
