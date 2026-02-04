package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l1rn/order-handler/internal/database"
	"github.com/l1rn/order-handler/internal/models"
)

type UserResponse struct {
	Username string `json:"username"`
}

func GetUsers(c *gin.Context){
	var users []models.User

	result := database.DB.Preload("Submission.Work").Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get all users"})
		return
	}
	var resp = make([]UserResponse, 0, len(users))
	
	for _, u := range users {
		resp = append(resp, UserResponse{
			Username: u.Username,
		})
	}

	c.JSON(http.StatusOK, resp)
}