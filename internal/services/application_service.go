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
	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var applicant models.Applicant
	if err := tx.First(&applicant, "id = ?", applicantID).Error; err != nil {
		tx.Rollback()
		return errors.New("applicant not found")
	}

	var scheme models.Scheme
	if err := tx.First(&scheme, "id = ?", schemeID).Error; err != nil {
		tx.Rollback()
		return errors.New("scheme not found")
	}

	var existingApplication models.Application
	if err := tx.Where("applicant_id = ? AND scheme_id = ?", applicantID, schemeID).
		First(&existingApplication).Error; err == nil {
		tx.Rollback()
		return errors.New("application already exists")
	}

	application := models.Application{
		ID:          utils.GenerateUUID(),
		ApplicantID: applicantID,
		SchemeID:    schemeID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := tx.Create(&application).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var application models.Application

	if err := tx.First(&application, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return errors.New("application not found")
	}

	application.ApplicantID = updatedData.ApplicantID
	application.SchemeID = updatedData.SchemeID
	application.UpdatedAt = time.Now()

	if err := tx.Save(&application).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DELETE Application
func (s *ApplicationService) DeleteApplication(applicationID string) error {
	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var application models.Application

	if err := tx.First(&application, "id = ?", applicationID).Error; err != nil {
		tx.Rollback()
		return errors.New("application not found")
	}

	if err := tx.Delete(&application).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DELETE Application by Applicant ID
func (s *ApplicationService) DeleteApplicationByApplicantID(applicantID string) error {
	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("applicant_id = ?", applicantID).Delete(&models.Application{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
