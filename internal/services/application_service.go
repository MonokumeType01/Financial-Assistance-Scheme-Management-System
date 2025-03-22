package services

import (
	"errors"
	"time"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/utils"
	"gorm.io/gorm"
)

type ApplicationService struct {
	DB *gorm.DB
}

func NewApplicationService(db *gorm.DB) *ApplicationService {
	return &ApplicationService{DB: db}
}

// CREATE Application
func (s *ApplicationService) RegisterApplication(applicantID, schemeID string) error {
	application := models.Application{
		ID:          utils.GenerateUUID(),
		ApplicantID: applicantID,
		SchemeID:    schemeID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.DB.Create(&application).Error
}

// RETRIEVE Applicantion by Applicant ID or Scheme ID
func (s *ApplicationService) GetApplications(applicantID, schemeID string) ([]models.Application, error) {
	var applications []models.Application
	query := s.DB

	if applicantID != "" {
		query = query.Where("applicant_id = ?", applicantID)
	}

	if schemeID != "" {
		query = query.Where("scheme_id = ?", schemeID)
	}

	if err := query.Find(&applications).Error; err != nil {
		return nil, err
	}

	return applications, nil
}

// RETRIEVE All Applications
func (s *ApplicationService) GetAllApplications() ([]models.Application, error) {
	var applications []models.Application
	if err := s.DB.Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

// UPDATE Application by ID
func (s *ApplicationService) UpdateApplication(id string, updatedData *models.Application) error {
	var application models.Application

	if err := s.DB.First(&application, "id = ?", id).Error; err != nil {
		return errors.New("application not found")
	}

	application.ApplicantID = updatedData.ApplicantID
	application.SchemeID = updatedData.SchemeID
	application.UpdatedAt = time.Now()

	return s.DB.Save(&application).Error
}

// DELETE Application
func (s *ApplicationService) DeleteApplication(applicationID string) error {
	var application models.Application

	// Check if application exists
	if err := s.DB.First(&application, "id = ?", applicationID).Error; err != nil {
		return errors.New("application not found")
	}

	// Perform delete
	return s.DB.Delete(&application).Error
}

// DELETE Application by Applicant ID
func (s *ApplicationService) DeleteApplicationByApplicantID(applicantID string) error {
	return s.DB.Where("applicant_id = ?", applicantID).Delete(&models.Application{}).Error
}
