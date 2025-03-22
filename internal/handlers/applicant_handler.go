package handlers

import (
	"net/http"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/dto"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/services"

	"github.com/gin-gonic/gin"
)

type ApplicantHandler struct {
	Service *services.ApplicantService
}

func NewApplicantHandler(service *services.ApplicantService) *ApplicantHandler {
	return &ApplicantHandler{Service: service}
}

// CREATE Applicant with Household
func (h *ApplicantHandler) CreateApplicant(c *gin.Context) {
	var data dto.ApplicantWithHousehold

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create applicant",
			"details": err.Error()})
		return
	}

	if err := h.Service.RegisterApplicantWithHousehold(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register applicant",
			"details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Applicant registered successfully"})
}

// RETRIEVE Applicant with Household
func (h *ApplicantHandler) GetApplicant(c *gin.Context) {
	id := c.Param("id")
	applicant, err := h.Service.GetApplicantWithID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Applicant not found",
			"details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"applicant": applicant})
}

// RETRIEVE All Applicants
func (h *ApplicantHandler) GetAllApplicants(c *gin.Context) {
	applicants, err := h.Service.GetApplicants(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve applicants",
			"details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"applicants": applicants})
}

// UDPATE applicant by ID
func (h *ApplicantHandler) UpdateApplicant(c *gin.Context) {
	id := c.Param("id")

	var data dto.ApplicantWithHousehold
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input format",
			"details": err.Error(),
		})
		return
	}

	if err := h.Service.UpdateApplicant(id, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update applicant",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Applicant updated successfully"})
}

// DELETE Applicant By ID
func (h *ApplicantHandler) DeleteApplicant(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.DeleteApplicant(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete applicant",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Applicant deleted successfully",
	})
}
