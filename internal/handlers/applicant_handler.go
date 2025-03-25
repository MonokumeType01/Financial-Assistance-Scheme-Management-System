package handlers

import (
	"net/http"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
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
	var data models.ApplicantWithHousehold

	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to create applicant")
		return
	}

	if err := h.Service.RegisterApplicantWithHousehold(&data); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to register applicant")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Applicant registered successfully"})
}

// RETRIEVE Applicant with Household
func (h *ApplicantHandler) GetApplicant(c *gin.Context) {
	id := c.Param("id")
	applicant, err := h.Service.GetApplicantWithID(id)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Applicant not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"applicant": applicant})
}

// RETRIEVE All Applicants
func (h *ApplicantHandler) GetAllApplicants(c *gin.Context) {
	applicants, err := h.Service.GetApplicants(c)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to retrieve applicants")
		return
	}

	c.JSON(http.StatusOK, gin.H{"applicants": applicants})
}

// UDPATE applicant by ID
func (h *ApplicantHandler) UpdateApplicant(c *gin.Context) {
	id := c.Param("id")

	var data models.ApplicantWithHousehold
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Invalid input format")
		return
	}

	if err := h.Service.UpdateApplicant(id, &data); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to update applicant")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Applicant updated successfully"})
}

// DELETE Applicant By ID
func (h *ApplicantHandler) DeleteApplicant(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.DeleteApplicant(id); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to delete applicant")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Applicant deleted successfully",
	})
}
