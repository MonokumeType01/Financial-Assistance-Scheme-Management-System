package routes

import (
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, applicantHandler *handlers.ApplicantHandler) {
	router.POST("/api/applicants", applicantHandler.CreateApplicant)
	router.GET("/api/applicants/:id", applicantHandler.GetApplicant)
}
