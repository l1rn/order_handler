package main

import (
	"net/http"

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

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/users", userController.GetUsers)
	router.Run()

}
