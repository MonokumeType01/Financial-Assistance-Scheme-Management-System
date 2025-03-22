package routes

import (
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, applicantHandler *handlers.ApplicantHandler, schemeHandler *handlers.SchemeHandler, applicationHandler *handlers.ApplicationHandler) {
	api := router.Group("/api")

	// Applicant
	applicantRoutes := api.Group("/applicants")
	{
		applicantRoutes.POST("/", applicantHandler.CreateApplicant)
		applicantRoutes.GET("/", applicantHandler.GetAllApplicants)
		applicantRoutes.GET("/:id", applicantHandler.GetApplicant)
		applicantRoutes.PUT("/:id", applicantHandler.UpdateApplicant)
		applicantRoutes.DELETE("/:id", applicantHandler.DeleteApplicant)
	}

	// Scheme
	schemeRoutes := api.Group("/schemes")
	{
		schemeRoutes.POST("/", schemeHandler.CreateScheme)
		schemeRoutes.GET("/", schemeHandler.GetAllSchemes)
		schemeRoutes.GET("/:id", schemeHandler.GetSchemeByID)
		schemeRoutes.PUT("/:id", schemeHandler.UpdateScheme)
		schemeRoutes.DELETE("/:id", schemeHandler.DeleteScheme)
		schemeRoutes.GET("/eligible/:applicantID", schemeHandler.GetEligibleSchemes)
	}

	// Applications
	applicationRoutes := api.Group("/applications")
	{
		applicationRoutes.POST("/", applicationHandler.RegisterApplication)
		applicationRoutes.GET("/", applicationHandler.GetApplications)
		applicationRoutes.PUT("/:id", applicationHandler.UpdateApplication)
		applicationRoutes.DELETE("/:id", applicationHandler.DeleteApplication)
		applicationRoutes.DELETE("/applicant/:applicant_id", applicationHandler.DeleteApplicationByApplicantID)
	}
}
