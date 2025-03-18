package handlers

import (
	"net/http"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/services"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/utils"

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
	var data models.ApplicantWithHouseHold

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.ID = utils.GenerateUUID()
	for i := range data.Household {
		data.Household[i].ID = utils.GenerateUUID()
		data.Household[i].ApplicantID = data.ID
	}

	if err := h.Service.RegisterApplicantWithHousehold(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register applicant"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Applicant registered successfully", "applicant": data})
}

// GET Applicant with Household
func (h *ApplicantHandler) GetApplicant(c *gin.Context) {
	id := c.Param("id")
	applicant, err := h.Service.GetApplicantWithHousehold(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Applicant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"applicant": applicant})
}
