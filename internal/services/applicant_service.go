package services

import (
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/utils"
	"gorm.io/gorm"
)

type ApplicantService struct {
	DB *gorm.DB
}

func NewApplicantService(db *gorm.DB) *ApplicantService {
	return &ApplicantService{DB: db}
}

// CREATE Applicant with Household Members
func (s *ApplicantService) RegisterApplicantWithHousehold(data *models.ApplicantWithHouseHold) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data.Applicant).Error; err != nil {
			return err
		}

		for i := range data.Household {
			data.Household[i].ID = utils.GenerateUUID()
			data.Household[i].ApplicantID = data.ID
		}

		if len(data.Household) > 0 {
			if err := tx.Create(&data.Household).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// RETRIEVE Applicant with Household Members
func (s *ApplicantService) GetApplicantWithHousehold(id string) (*models.ApplicantWithHouseHold, error) {
	var applicant models.Applicant
	var household []models.HouseholdMember

	if err := s.DB.First(&applicant, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Find(&household, "applicant_id = ?", id).Error; err != nil {
		return nil, err
	}

	return &models.ApplicantWithHouseHold{
		Applicant: applicant,
		Household: household,
	}, nil
}
