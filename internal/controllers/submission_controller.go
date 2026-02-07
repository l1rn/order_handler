package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/l1rn/order-handler/internal/services"
)

type SubmissionController struct {
	submissionService services.SubmissionService
}

func NewSubmissionController(s services.SubmissionService) *SubmissionController {
	return &SubmissionController{submissionService: s}
}

func (ctrl *SubmissionController) GetSubmissions(c *gin.Context) {
	subs, err := ctrl.submissionService.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to find all submissions",
			"details": err.Error(),
		})
		return
	}

	c.JSON(200, subs)
}

func (ctrl *SubmissionController) AddWorkItem(c *gin.Context) {
	idP := c.Query("id")
	wiP := c.Query("wi")
	if idP == "" || wiP == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id and wi are required"})
		return
	}

	id, err1 := strconv.ParseUint(idP, 10, 32)
	wi, err2 := strconv.ParseUint(wiP, 10, 32)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id and wi must be numbers"})
		return
	}
	err := ctrl.submissionService.AddWorkItemToSubmission(uint(id), uint(wi))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "work item added successfully"})
}
