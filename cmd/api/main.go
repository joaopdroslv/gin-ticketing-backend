package main

import (
	"context"
	"log"

	"go-gin-ticketing-backend/internal/config"
	"go-gin-ticketing-backend/internal/database"

	authmiddleware "go-gin-ticketing-backend/internal/auth/middleware"

	authhandler "go-gin-ticketing-backend/internal/auth/handler"
	userhandler "go-gin-ticketing-backend/internal/user/handler"

	authservice "go-gin-ticketing-backend/internal/auth/service"
	userservice "go-gin-ticketing-backend/internal/user/service"

	authrepository "go-gin-ticketing-backend/internal/auth/repository"
	userrepository "go-gin-ticketing-backend/internal/user/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	env := config.NewEnv()
	logger := config.NewLogger()

	db, err := database.NewMysql(env.DockerDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	// r.Use(gin.Logger(), gin.Recovery())

	apiV1Group := r.Group("/api/v1")

	// repositories
	userRepo := userrepository.NewUserRepositoryMysql(db)
	userStatusRepo := userrepository.NewUserStatusRepositoryMysql(db)
	authRepo := authrepository.NewAuthRepositoryMysql(db)
	permissionRepo := authrepository.NewPermissionRepositoryMysql(db)

	// services
	userStatusService := userservice.NewUserStatusService(userStatusRepo)
	ctx := context.Background()
	userService, err := userservice.NewUserService(ctx, userRepo, userStatusService)
	if err != nil {
		log.Fatal("failed to create the user service")
	}
	authService := authservice.New(authRepo, permissionRepo, env.JWTSecret, env.JWTTTL)

	// handlers
	authHandler := authhandler.New(authService)
	userHandler := userhandler.New(logger, userService)

	// middlewares
	jwtMiddleware := authmiddleware.JWTAuthenticationMiddleware(env.JWTSecret)

	// routes

	// auth (public)
	authGroup := apiV1Group.Group("/auth")
	{
		authGroup.POST("/register", authHandler.RegisterUser)
		authGroup.POST("/login", authHandler.LoginUser)
	}

	// users (protected)
	userGroup := apiV1Group.Group("/users")
	userGroup.Use(jwtMiddleware)
	userhandler.RegisterRoutes(userGroup, userHandler, authService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Ok"})
	})

	r.Run(":8080")
}
