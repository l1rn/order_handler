package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l1rn/order-handler/internal/database"
	"github.com/l1rn/order-handler/internal/models"
)

func GetUsers(c *gin.Context){
	var users []models.User

	result := database.DB.Preload("Submission.Work").Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get all users"})
		return
	}

	c.JSON(http.StatusOK, users)
}