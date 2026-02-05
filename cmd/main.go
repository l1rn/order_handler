package main

import (
	"github.com/gin-gonic/gin"
	"github.com/l1rn/order-handler/internal/controllers"
	"github.com/l1rn/order-handler/internal/database"
	"github.com/l1rn/order-handler/internal/repositories"
	"github.com/l1rn/order-handler/internal/services"
)

func main() {
	db := database.InitDB()
	database.InitMockData(db)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	workRepo := repositories.NewWorkRepository(db)
	workService := services.NewWorkService(workRepo)
	workController := controllers.NewWorkController(workService)

	r := gin.Default()

	userRoutes := r.Group("/api/v1/users")

	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.GetUsers)
		userRoutes.GET("/:id", userController.GetUserById)
	}

	workRoutes := r.Group("/api/v1/work-items")

	{
		workRoutes.GET("", workController.GetAll)
		workRoutes.POST("", workController.CreateWorkItem)
	}
	r.Run()
}
