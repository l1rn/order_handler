package main

import (
	"github.com/gin-gonic/gin"
	"github.com/l1rn/order-handler/internal/controllers"
	"github.com/l1rn/order-handler/internal/database"
	"github.com/l1rn/order-handler/internal/repository"
	"github.com/l1rn/order-handler/internal/service"
)

func main() {
	db := database.InitDB()
	database.InitMockData(db)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	userRoutes := r.Group("/api/v1")

	{
		userRoutes.GET("/users", userController.GetUsers)
		userRoutes.GET("/users/:id", userController.GetUserById)
	}

	r.Run()

}
