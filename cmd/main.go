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
	database.SeedData(db)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	workRepo := repositories.NewWorkRepository(db)
	workService := services.NewWorkService(workRepo)
	workController := controllers.NewWorkController(workService)

	submissionRepo := repositories.NewSubmissionRepository(db)
	submissionService := services.NewSubmissionService(submissionRepo)
	submissionController := controllers.NewSubmissionController(submissionService)

	authController := controllers.NewAuthController(userService)

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
		workRoutes.PUT("/:id", workController.UpdateWorkItem)
	}

	subRoutes := r.Group("/api/v1/submissions")

	{
		subRoutes.GET("", submissionController.GetSubmissions)
		subRoutes.POST("/add-work/", submissionController.AddWorkItem)
	}

	authRoutes := r.Group("/api/v1/auth")
	{
		authRoutes.POST("/sign-up", authController.Register)
	}
	r.Run(":8081")
}
