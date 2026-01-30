package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l1rn/order-handler/internal/controllers"
	"github.com/l1rn/order-handler/internal/database"
)

func main() {
	database.InitDB()
	database.InitMockData()
	
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/users", controllers.GetUsers)
	router.Run()
}
