// Package modules: main modules to be used as an entry point
package modules

import (
	blog_usecase "blogger/modules/blog/application/usecases"
	blog_routes "blogger/modules/blog/delivery/http/routes"
	blog_repository "blogger/modules/blog/infrastructure/persistence/repository"
	user_usecase "blogger/modules/user/application/usecases"
	user_routes "blogger/modules/user/delivery/http/routes"
	user_repository "blogger/modules/user/infrastructure/persistence/repository"
	"blogger/modules/utils"
	"blogger/modules/utils/middleware"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type BloggerServer struct {
	eng *gin.Engine
	client *mongo.Client
	db *mongo.Database
	cfg utils.Config
}

func NewBloggerServer() *BloggerServer {
	cfg := utils.NewConfig()
	client, err := utils.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Mongo database connection failed %s\n", err.Error())
	}

	db := client.Database(cfg.DatabaseName)

	return &BloggerServer{
		eng: gin.Default(),
		cfg: cfg,
		client: client,
		db: db,
	}
}

func handlePing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func (b *BloggerServer) prepare() {
	apiSubGroup := b.eng.Group("/api/v1")
	// prepare user endpoints
	userCollection:= utils.CreateCollection(b.db, "userCollection")
	blogCollection := utils.CreateCollection(b.db, "blogCollection")
	tokenCollection := utils.CreateCollection(b.db, "tokenCollection")
	
	tokenUtil := utils.NewJwtService(b.cfg.SecretKey)

	
	userRepository := user_repository.NewUserRepository(userCollection)
	tokenRepository := user_repository.NewTokenRepository(tokenCollection)
	blogRepository := blog_repository.NewBlogRepository(blogCollection)
	
	authMiddleware := middleware.JwtAuthMiddleware(userRepository, tokenUtil)
	
	userUseCase := user_usecase.NewUserUseCase(userRepository, tokenRepository, utils.NewPasswordUtil(), tokenUtil)
	userRoutes := user_routes.NewUserRoutes(userUseCase, authMiddleware)
	userRoutes.RegisterRoutes(apiSubGroup)
	
	blogUseCase := blog_usecase.NewBlogUseCase(blogRepository)
	blogRoutes := blog_routes.NewBlogRoutes(blogUseCase, authMiddleware)
	blogRoutes.RegisterRoutes(apiSubGroup)

	// prepare blog endpoints
	b.eng.GET("/", handlePing)
}

func (b *BloggerServer) disconnect(ctx context.Context) {
	b.client.Disconnect(ctx)
}

func (b *BloggerServer)RunServer() {
	b.prepare()

	b.eng.Run(b.cfg.Port)

	// srv := http.Server{
	// 	Addr: b.cfg.Port,
	// 	Handler: b.eng,
	// }

	// go func() {
	// 	log.Printf("Server running on http://localhost:%s\n", b.cfg.Port)
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("Listen server error: %v\n", err.Error())
	// 	}
	// }()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Println("Shutdown signal is recived")

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	// defer cancel()	

	// b.disconnect(ctx)
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Forced server shutdown")
	// }

	// log.Println("Server was shutdown gracefully");
}

