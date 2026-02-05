package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/l1rn/order-handler/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{userService: s}
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.userService.FindAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUserById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID Format"})
		return
	}

	user, err := ctrl.userService.FindById(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to find user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
