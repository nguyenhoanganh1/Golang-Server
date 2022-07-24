package main

import (
	"example/web-service-gin/config"
	"example/web-service-gin/controller"
	"example/web-service-gin/repository"
	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB = config.SetupDatabaseSQLConnection()
	userRepository          = repository.NewUserRepository(db)
	jwtService              = service.NewJWTService()
	authService             = service.NewAuthService(userRepository)
	authController          = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := gin.Default()
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	router.Run("localhost:8080")
}
