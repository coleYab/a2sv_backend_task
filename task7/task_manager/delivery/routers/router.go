// Package routes: routes module
package routes

import (
	"context"
	"log"
	"os"
	"task_manager/delivery/controllers"
	"task_manager/infrastructure"
	"task_manager/repositories"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI string = os.Getenv("MONGO_URI")
var secretKey string = os.Getenv("SECRET_KEY")

type Router struct {
	client         *mongo.Client
	userRepository repositories.IUserRepository
	taskController controllers.ITaskController
	userController controllers.IUserController
	jwtService     infrastructure.IJwtService
	e              *gin.Engine
}

func New() *Router {
	opts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("unable to connect to mongo db: %v\n", err.Error())
	}

	taskCollection := client.Database("db").Collection("task_collection")
	userCollection := client.Database("db").Collection("user_collection")
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("unable to ping mongo db: %v\n", err.Error())
	}

	log.Println("connected to mongo db")
	taskRepository := repositories.NewTaskRepository(taskCollection)
	taskController := controllers.NewTaskController(taskRepository)

	// for the auth
	userRepository := repositories.NewUserRepository(userCollection)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUseCase, infrastructure.NewJwtService(secretKey))

	return &Router{
		client:         client,
		userRepository: userRepository,
		taskController: taskController,
		userController: userController,
		jwtService:     infrastructure.NewJwtService(secretKey),
		e:              gin.Default(),
	}
}

func (r *Router) Run(addr string) {
	// it will always throw an error when the server is closed
	if err := r.e.Run(addr); err != nil {
		log.Println("Server stopped!")
		r.client.Disconnect(context.TODO())
		log.Println("Database disconnected!!")
	}
}

func (r *Router) RegisterRoutes() {
	// this is the task controllers
	taskRoutes := r.e.Group("/task")
	taskRoutes.Use(infrastructure.JwtAuthMiddleware(r.userRepository, r.jwtService))
	{
		taskRoutes.GET("/", r.taskController.GetTasks)
		taskRoutes.GET("/:id", r.taskController.GetTaskByID)
		taskRoutes.PUT("/:id", r.taskController.UpdateTask)
		taskRoutes.DELETE("/:id", r.taskController.DeleteTask)
		taskRoutes.POST("/", r.taskController.CreateTask)
	}

	// this is the auth controllers
	r.e.POST("/register", r.userController.RegisterUser)
	r.e.POST("/login", r.userController.Login)
	r.e.POST("/user/:id/promote", infrastructure.JwtAuthMiddleware(r.userRepository, r.jwtService), r.userController.PromoteUser)
}
