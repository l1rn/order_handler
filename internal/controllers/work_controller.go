package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l1rn/order-handler/internal/models"
	"github.com/l1rn/order-handler/internal/services"
)

type WorkController struct {
	workService services.WorkService
}

func NewWorkController(s services.WorkService) *WorkController {
	return &WorkController{workService: s}
}

func (ctrl *WorkController) GetAll(c *gin.Context) {
	works, err := ctrl.workService.FindAllWorkItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch work items",
			"details": err.Error(),
		})
	}

	c.JSON(200, works)
}

func (ctrl *WorkController) CreateWorkItem(c *gin.Context) {
	var req models.CreateWorkItemRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	err := ctrl.workService.CreateWorkItem(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "work item was created!"})
}
