package router

import (
	"context"
	"log"
	"os"
	"task_manager/controllers"
	"task_manager/data"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUri string = os.Getenv("MONGO_URI")

type Router struct {
	client *mongo.Client
	c controllers.ITaskController
	e *gin.Engine
}

func New() *Router {
	opts := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("unable to connect to mongo db: %v\n", err.Error())
	}

	taskCollection := client.Database("db").Collection("task_collection")
	
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("unable to ping mongo db: %v\n", err.Error())
	}

	log.Println("connected to mongo db")
	taskService := data.New(taskCollection)
	taskController := controllers.New(taskService)

	return &Router{
		client: client,
		c: taskController,
		e: gin.Default(),
	}
}

func (r *Router)Run(addr string) {
	// it will always throw an error when the server is closed 
	if err := r.e.Run(addr); err != nil {
		log.Println("Server stopped!")
		r.client.Disconnect(context.TODO())
		log.Println("Database disconnected!!")
	}
}

func (r *Router)RegisterRoutes() {
	r.e.GET("/task", r.c.GetTasks)
	r.e.GET("/task/:id", r.c.GetTaskById)
	r.e.PUT("/task/:id", r.c.UpdateTask)
	r.e.DELETE("/task/:id", r.c.DeleteTask)
	r.e.POST("/task", r.c.CreateTask)
}