package data

import (
	"context"
	"fmt"
	"task_manager/controllers/dto"
	"task_manager/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ITaskService interface {
	GetTasks() []models.Task
	DeleteTask(id string) error
	GetTaskById(id string) (models.Task, error)
	CreateTask(task dto.CreateTaskDto) models.Task
	UpdateTask(id string, task dto.UpdateTaskDto) error
}

type TaskService struct {
	collection *mongo.Collection
	tasks map[string]models.Task
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{
		collection: collection,
		tasks: map[string]models.Task{},
	}
}

func (ts *TaskService) GetTasks() []models.Task {
	findOpts := options.Find()
	cursor, err := ts.collection.Find(context.TODO(), bson.D{{}}, findOpts)
	if err != nil {
		return nil
	}

	tasks := []models.Task{}

	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			// returning err
		}
		tasks = append(tasks, task)
	}

	cursor.Close(context.TODO())
	return tasks
}

func (ts *TaskService) GetTaskById(id string) (models.Task, error) {
	findOneOpts := options.FindOne()	
	filter := bson.D{{Key: "id", Value: id}}

	var task models.Task
	if err := ts.collection.FindOne(context.TODO(), filter, findOneOpts).Decode(&task); err != nil {
		return models.Task{}, fmt.Errorf("find failed %v", err.Error()) 
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

	ts.collection.InsertOne(context.TODO(), newTask)

	return newTask
}

func (ts *TaskService)DeleteTask(id string) error {
	findOneAndDeleteOpts := options.FindOneAndDelete()
	filters := bson.D{{Key: "id", Value: id}}
	var task models.Task
	if err := ts.collection.FindOneAndDelete(context.TODO(), filters, findOneAndDeleteOpts).Decode(&task); err != nil {
		return err
	}

	return nil
}

func (ts *TaskService)UpdateTask(id string, task dto.UpdateTaskDto) error {
	oldtask, err := ts.GetTaskById(id)
	if err != nil {
		return err
	}

	filter := bson.D{{ Key: "id", Value: id}}
	findOneOpts := options.FindOneAndReplace()

	oldtask.Title = task.Title
	oldtask.Description = task.Description
	oldtask.DueDate = task.DueDate
	oldtask.Status = task.Status

	ts.collection.FindOneAndReplace(context.TODO(), filter, oldtask, findOneOpts)

	return nil
}
