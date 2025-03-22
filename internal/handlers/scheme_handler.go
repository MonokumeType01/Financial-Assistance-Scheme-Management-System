package handlers

import (
	"net/http"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/services"

	"github.com/gin-gonic/gin"
)

type SchemeHandler struct {
	Service *services.SchemeService
}

func NewSchemeHandler(service *services.SchemeService) *SchemeHandler {
	return &SchemeHandler{Service: service}
}

// CREATE Scheme
func (h *SchemeHandler) CreateScheme(c *gin.Context) {
	var data models.Scheme
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Invalid input format")
		return
	}

	if err := h.Service.CreateScheme(&data); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to create scheme")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Scheme created successfully"})
}

// RETRIEVE All Scheme
func (h *SchemeHandler) GetAllSchemes(c *gin.Context) {
	schemes, err := h.Service.GetAllSchemes()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to retrieve schemes")
	}

	c.JSON(http.StatusOK, gin.H{"schemes": schemes})
}

// RETRIEVE Scheme by ID
func (h *SchemeHandler) GetSchemeByID(c *gin.Context) {
	id := c.Param("id")
	scheme, err := h.Service.GetSchemeByID(id)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Scheme not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"scheme": scheme})
}

// UPDATE Scheme by ID
func (h *SchemeHandler) UpdateScheme(c *gin.Context) {
	id := c.Param("id")

	var updatedData models.Scheme
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Invalid input data")
		return
	}

	if err := h.Service.UpdateScheme(id, &updatedData); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to update scheme")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Scheme updated successfully"})
}

// DELETE Scheme by ID
func (h *SchemeHandler) DeleteScheme(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.DeleteScheme(id); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to delete scheme")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Scheme deleted successfully"})
}

// RETRIEVE Eligible Schemes
func (h *SchemeHandler) GetEligibleSchemes(c *gin.Context) {
	applicantID := c.Param("applicantID")

	eligibleSchemes, err := h.Service.GetEligibleSchemes(applicantID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta("Failed to get eligible scheme")
		return
	}

	c.JSON(http.StatusOK, gin.H{"eligible_schemes": eligibleSchemes})
}
