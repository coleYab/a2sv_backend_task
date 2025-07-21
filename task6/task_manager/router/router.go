package router

import (
	"context"
	"log"
	"os"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/middleware"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUri string = os.Getenv("MONGO_URI")

func ProtectRoute(userStore models.User, ctx *gin.Context) func (*gin.Context) {
	return func (ctx *gin.Context)  {
		// here do the magic and add the user to the context
	}
}

type Router struct {
	client *mongo.Client
	userService data.IUserService
	taskController controllers.ITaskController
	userController controllers.IUserController	
	e *gin.Engine
}

func New() *Router {
	opts := options.Client().ApplyURI(mongoUri)
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
	taskService := data.NewTaskService(taskCollection)
	taskController := controllers.NewTaskController(taskService)

	// for the auth
	userService := data.NewUserService(userCollection)
	userController := controllers.NewUserController(userService)

	return &Router{
		client: client,
		userService: userService,
		taskController: taskController,
		userController: userController,
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
	// this is the task controllers 
	taskRoutes := r.e.Group("/task")
	taskRoutes.Use(middleware.JwtAuthMiddleware(r.userService)) 
	{
		taskRoutes.GET("/", r.taskController.GetTasks)
		taskRoutes.GET("/:id", r.taskController.GetTaskById)
		taskRoutes.PUT("/:id", r.taskController.UpdateTask)
		taskRoutes.DELETE("/:id", r.taskController.DeleteTask)
		taskRoutes.POST("/", r.taskController.CreateTask)
	}

	// this is the auth controllers
	r.e.POST("/register", r.userController.RegisterUser)
	r.e.POST("/login", r.userController.Login)
	r.e.POST("/user/:id/promote", middleware.JwtAuthMiddleware(r.userService), r.userController.PromoteUser)
}