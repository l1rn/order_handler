package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/l1rn/order-handler/internal/models"
	"github.com/l1rn/order-handler/internal/services"
)

type AuthorizationController struct {
	userService services.UserService
}

func NewAuthController(us services.UserService) *AuthorizationController {
	return &AuthorizationController{userService: us}
}

var jwtKey = []byte("test_key")

func generateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (ctrl *AuthorizationController) Register(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request", "details": err.Error()})
		return
	}
	userId, err := ctrl.userService.CreateUser(req)
	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to create user", "detailes": err.Error()})
		return
	}

	token, _ := generateJWT(userId)

	c.JSON(201, gin.H{"token": token})
}
