package main

import (
	"context"
	"log"
	"time"

	"go-gin-ticketing-backend/internal/config"
	"go-gin-ticketing-backend/internal/database"
	"go-gin-ticketing-backend/internal/middlewares"

	authhandler "go-gin-ticketing-backend/internal/auth/handler"
	userhandler "go-gin-ticketing-backend/internal/user/handler"

	accesscontrolservice "go-gin-ticketing-backend/internal/access_control/service"
	authservice "go-gin-ticketing-backend/internal/auth/service"
	userservice "go-gin-ticketing-backend/internal/user/service"

	accesscontrolrepository "go-gin-ticketing-backend/internal/access_control/repository"
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

	// middlewares
	jwtMiddleware := middlewares.JWTAuthenticationMiddleware(env.JWTSecret)
	rateLimitMiddleware := middlewares.RateLimitMiddleware(env.RequestsPerMinute, time.Minute)

	// All routes covered by the rate limit middleware
	r.Use(rateLimitMiddleware)

	apiV1Group := r.Group("/api/v1")

	// repositories
	userRepo := userrepository.NewUserRepositoryMysql(db)
	userStatusRepo := userrepository.NewUserStatusRepositoryMysql(db)
	authRepo := authrepository.NewAuthRepositoryMysql(db)
	permissionRepo := accesscontrolrepository.NewPermissionRepositoryMysql(db)

	// services
	userStatusService := userservice.NewUserStatusService(userStatusRepo)
	ctx := context.Background()
	userService, err := userservice.NewUserService(ctx, userRepo, userStatusService)
	if err != nil {
		log.Fatal("failed to create the user service")
	}
	authService := authservice.New(authRepo, env.JWTSecret, env.JWTTTL)
	permissionService := accesscontrolservice.NewPermissionService(permissionRepo)

	// handlers
	authHandler := authhandler.New(authService)
	userHandler := userhandler.New(logger, userService)

	// routes

	// auth (public)
	authGroup := apiV1Group.Group("/auth")
	{
		authGroup.POST("/login", authHandler.LoginUser)
		authGroup.POST("/register", authHandler.RegisterUser)
	}

	// users (protected)
	userGroup := apiV1Group.Group("/users")
	userGroup.Use(jwtMiddleware)
	userhandler.RegisterRoutes(userGroup, userHandler, permissionService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Ok"})
	})

	r.Run(":" + env.HTTPPort)
}
