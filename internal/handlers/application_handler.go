package handlers

import (
	"net/http"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/services"
	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct {
	Service *services.ApplicationService
}

func NewApplicationHandler(service *services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{Service: service}
}

// CREATE Application
func (h *ApplicationHandler) RegisterApplication(c *gin.Context) {
	var input struct {
		ApplicantID string `json:"applicant_id"`
		SchemeID    string `json:"scheme_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	if err := h.Service.RegisterApplication(input.ApplicantID, input.SchemeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register application"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Application registered successfully"})
}

// RETRIEVE Application by Applicant ID or Scheme ID
func (h *ApplicationHandler) GetApplications(c *gin.Context) {
	applicantID := c.Query("applicant_id")
	schemeID := c.Query("scheme_id")

	applications, err := h.Service.GetApplications(applicantID, schemeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve applications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"applications": applications})
}

// RETRIEVE All Applications
func (h *ApplicationHandler) GetAllApplications(c *gin.Context) {
	applications, err := h.Service.GetAllApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve applications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"applications": applications})
}

// UPDATE Application
func (h *ApplicationHandler) UpdateApplication(c *gin.Context) {
	id := c.Param("id")

	var data models.Application
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	if err := h.Service.UpdateApplication(id, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update application"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application updated successfully"})
}

// DELETE Application
func (h *ApplicationHandler) DeleteApplication(c *gin.Context) {
	applicationID := c.Param("id")

	if err := h.Service.DeleteApplication(applicationID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Failed to delete application",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}

// DELETE Application by Applicant ID
func (h *ApplicationHandler) DeleteApplicationByApplicantID(c *gin.Context) {
	applicantID := c.Param("applicant_id")

	if err := h.Service.DeleteApplicationByApplicantID(applicantID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete applications"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Applications deleted successfully"})
}
