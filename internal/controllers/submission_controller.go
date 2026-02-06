package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l1rn/order-handler/internal/services"
)

type SubmissionController struct{
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