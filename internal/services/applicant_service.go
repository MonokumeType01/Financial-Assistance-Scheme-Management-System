package services

import (
	"context"
	"errors"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/dto"
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
func (s *ApplicantService) RegisterApplicantWithHousehold(data *dto.ApplicantWithHousehold) error {
	applicant := models.Applicant{
		ID:               utils.GenerateUUID(),
		Name:             data.Name,
		EmploymentStatus: data.EmploymentStatus,
		Sex:              data.Sex,
		DateOfBirth:      data.DateOfBirth,
	}

	householdMembers := make([]models.HouseholdMember, len(data.Household))
	for i, member := range data.Household {
		householdMembers[i] = models.HouseholdMember{
			ID:               utils.GenerateUUID(),
			Name:             member.Name,
			EmploymentStatus: member.EmploymentStatus,
			Sex:              member.Sex,
			DateOfBirth:      member.DateOfBirth,
			Relation:         member.Relation,
			SchoolLevel:      member.SchoolLevel,
			ApplicantID:      applicant.ID,
		}
	}

	return s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&applicant).Error; err != nil {
			return err
		}

		if len(householdMembers) > 0 {
			if err := tx.Create(&householdMembers).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// RETRIEVE All Applicant with Household Members
func (s *ApplicantService) GetApplicants(ctx context.Context) ([]dto.ApplicantWithHousehold, error) {
	var applicants []models.Applicant
	if err := s.DB.Preload("Household").Find(&applicants).Error; err != nil {
		return nil, err
	}

	output := make([]dto.ApplicantWithHousehold, len(applicants))
	for i, applicant := range applicants {
		householdDTO := make([]dto.HouseholdMember, len(applicant.Household))
		for j, member := range applicant.Household {
			householdDTO[j] = dto.HouseholdMember{
				ID:               member.ID,
				Name:             member.Name,
				EmploymentStatus: member.EmploymentStatus,
				Sex:              member.Sex,
				DateOfBirth:      member.DateOfBirth,
				Relation:         member.Relation,
			}
		}

		output[i] = dto.ApplicantWithHousehold{
			Applicant: dto.Applicant{
				ID:               applicant.ID,
				Name:             applicant.Name,
				EmploymentStatus: applicant.EmploymentStatus,
				Sex:              applicant.Sex,
				DateOfBirth:      applicant.DateOfBirth,
			},
			Household: householdDTO,
		}
	}

	return output, nil
}

// RETRIEVE Applicant with Household Members by Applicant ID
func (s *ApplicantService) GetApplicantWithID(id string) (*dto.ApplicantWithHousehold, error) {
	var applicant models.Applicant
	var household []models.HouseholdMember

	if err := s.DB.First(&applicant, "id = ?", id).Error; err != nil {
		return nil, errors.New("applicant not found")
	}

	if err := s.DB.Find(&household, "applicant_id = ?", id).Error; err != nil {
		return nil, errors.New("failed to retrieve household members")
	}

	householdDTO := make([]dto.HouseholdMember, len(household))
	for i, member := range household {
		householdDTO[i] = dto.HouseholdMember{
			ID:               member.ID,
			Name:             member.Name,
			EmploymentStatus: member.EmploymentStatus,
			Sex:              member.Sex,
			DateOfBirth:      member.DateOfBirth,
			Relation:         member.Relation,
		}
	}

	return &dto.ApplicantWithHousehold{
		Applicant: dto.Applicant{
			ID:               applicant.ID,
			Name:             applicant.Name,
			EmploymentStatus: applicant.EmploymentStatus,
			Sex:              applicant.Sex,
			DateOfBirth:      applicant.DateOfBirth,
		},
		Household: householdDTO,
	}, nil
}

// UDPATE applicant by ID
func (s *ApplicantService) UpdateApplicant(id string, updatedData *dto.ApplicantWithHousehold) error {
	var applicant models.Applicant

	if err := s.DB.First(&applicant, "id = ?", id).Error; err != nil {
		return errors.New("applicant not found")
	}

	return s.DB.Transaction(func(tx *gorm.DB) error {
		applicant.Name = updatedData.Name
		applicant.EmploymentStatus = updatedData.EmploymentStatus
		applicant.Sex = updatedData.Sex
		applicant.DateOfBirth = updatedData.DateOfBirth

		if err := tx.Save(&applicant).Error; err != nil {
			return err
		}

		if err := tx.Where("applicant_id = ?", applicant.ID).Delete(&models.HouseholdMember{}).Error; err != nil {
			return err
		}

		existingHouseholdIDs := make(map[string]bool)
		for _, existingMember := range applicant.Household {
			existingHouseholdIDs[existingMember.ID] = true
		}

		householdMembers := make([]models.HouseholdMember, len(updatedData.Household))
		for i, member := range updatedData.Household {
			if existingHouseholdIDs[member.ID] {
				householdMembers[i] = models.HouseholdMember{
					ID:               member.ID,
					Name:             member.Name,
					EmploymentStatus: member.EmploymentStatus,
					Sex:              member.Sex,
					DateOfBirth:      member.DateOfBirth,
					Relation:         member.Relation,
					SchoolLevel:      member.SchoolLevel,
					ApplicantID:      applicant.ID,
				}
			} else {
				householdMembers[i] = models.HouseholdMember{
					ID:               utils.GenerateUUID(),
					Name:             member.Name,
					EmploymentStatus: member.EmploymentStatus,
					Sex:              member.Sex,
					DateOfBirth:      member.DateOfBirth,
					Relation:         member.Relation,
					SchoolLevel:      member.SchoolLevel,
					ApplicantID:      applicant.ID,
				}
			}
		}

		if len(householdMembers) > 0 {
			if err := tx.Create(&householdMembers).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// DELETE Applicant By ID
func (s *ApplicantService) DeleteApplicant(id string) error {
	var applicant models.Applicant
	if err := s.DB.First(&applicant, "id = ?", id).Error; err != nil {
		return errors.New("applicant not found")
	}

	if err := s.DB.Where("applicant_id = ?", id).Delete(&models.HouseholdMember{}).Error; err != nil {
		return errors.New("failed to delete household members")
	}

	// TODO: Add after applicant endpoints are done
	// if err := s.DB.Where("applicant_id = ?", id).Delete(&models.Application{}).Error; err != nil {
	// 	return errors.New("failed to delete applicant's application")
	// }

	if err := s.DB.Delete(&applicant).Error; err != nil {
		return errors.New("failed to delete applicant")
	}

	return nil
}
