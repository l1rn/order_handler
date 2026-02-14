package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/l1rn/order-handler/internal/models"
	"github.com/l1rn/order-handler/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthorizationController struct {
	userService services.UserService
}

func NewAuthController(us services.UserService) *AuthorizationController {
	return &AuthorizationController{userService: us}
}

var jwtKey = []byte("test_key")

func generateTokenPair(userID uint) (string, string) {
	accessTClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTClaims)
	accessToken, _ := at.SignedString(jwtKey)

	refreshTClaims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTClaims)
	refreshToken, _ := rt.SignedString(jwtKey)

	return accessToken, refreshToken
}

func (ctrl *AuthorizationController) Register(c *gin.Context) {
	var req models.UserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request", 
			"details": err.Error(),
		})
		return
	}

	user, err := ctrl.userService.FindByUsername(req.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"details": err.Error(),
		})
		return
	}
	if user != nil {
		c.JSON(409, gin.H{
			"message": "The username already exists", 
		})
		return
	}

	_, err = ctrl.userService.CreateUser(req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create user", 
			"detailes": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"message": "User was created!"})
}

func (ctrl *AuthorizationController) Login(c *gin.Context) {
	var loginRequest models.UserRequest
	
	if err := c.ShouldBindBodyWithJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to login",
			"details": err.Error(),
		})
		return
	}

	user, err := ctrl.userService.FindByUsername(loginRequest.Username)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid Credentials",
			"details": err.Error(),
		})
		return 
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password), 
		[]byte(loginRequest.Password),
	)

	if err != nil {
		c.JSON(401, gin.H {
			"message": "Invalid Credentials",
			"details": err.Error(),
		})
		return
	}

	at, rt := generateTokenPair(user.ID)
	
	c.SetCookie("access_token", at, 15, "/", "", true, true)
	c.SetCookie("refresh_token", rt, 3600*24*7, "/", "", true, true)
} 

func (ctrl *AuthorizationController) Refresh(c *gin.Context){
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Refresh token missing",
			"details": err.Error(),
		})
	}

	token, _ := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		newAccess, newRefresh := generateTokenPair(userID)
		c.SetCookie("access_token", newAccess, 15, "/", "", true, true)
		c.SetCookie("refresh_token", newRefresh, 3600*24*7, "/", "", true, true)
	}
}

func ProxyAuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		userID := c.GetHeader("X-User-ID")

		if userID == ""{
			c.AbortWithStatusJSON(401, gin.H{"error": "Internal Auth Error"})
			return
		}

		c.Set("currentUser", userID)
		c.Next()
	}
}