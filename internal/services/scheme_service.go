package services

import (
	"errors"
	"log"
	"time"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/data"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/dto"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/utils"
	"gorm.io/gorm"
)

type SchemeService struct {
	DB *gorm.DB
}

func NewSchemeService(db *gorm.DB) *SchemeService {
	return &SchemeService{DB: db}
}

/* Helper Functions */
func isEligible(applicant models.Applicant, criteria models.Criteria) bool {
	if criteria.EmploymentStatus != "" && applicant.EmploymentStatus != criteria.EmploymentStatus {
		return false
	}

	if criteria.HasChildren != nil {
		hasEligibleChild := false
		for _, householdMember := range applicant.Household {
			if householdMember.Relation == data.RELATION_SON || householdMember.Relation == data.RELATION_DAUGHTER {

				childLevel := householdMember.SchoolLevel
				schemeLevel := criteria.HasChildren.SchoolLevel
				condition := criteria.HasChildren.SchoolLevelCondition

				switch condition {
				case data.CRITERIA_EQUAL:
					if childLevel == schemeLevel {
						hasEligibleChild = true
					}
				case data.CRITERIA_EQUAL_OR_ABOVE:
					if childLevel >= schemeLevel {
						hasEligibleChild = true
					}
				case data.CRITERIA_EQUAL_OR_BELOW:
					if childLevel <= schemeLevel {
						hasEligibleChild = true
					}
				case data.CRITERIA_ABOVE:
					if childLevel > schemeLevel {
						hasEligibleChild = true
					}
				case data.CRITERIA_BELOW:
					if childLevel < schemeLevel {
						hasEligibleChild = true
					}
				default:
					log.Printf("[ERROR] Unknown condition: %v", condition)
				}

				if hasEligibleChild {
					break
				}
			}
		}

		if !hasEligibleChild {
			return false
		}
	}

	return true
}

/* Service Functions */

// CREATE Scheme
func (s *SchemeService) CreateScheme(schemeData *models.Scheme) error {
	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := utils.ValidateScheme(
		schemeData.Name,
		schemeData.Criteria.EmploymentStatus,
		schemeData.Criteria.HasChildren,
	); err != nil {
		return err
	}

	scheme := models.Scheme{
		ID:   utils.GenerateUUID(),
		Name: schemeData.Name,
		Criteria: models.Criteria{
			EmploymentStatus: schemeData.Criteria.EmploymentStatus,
			HasChildren: &models.Children{
				SchoolLevel:          schemeData.Criteria.HasChildren.SchoolLevel,
				SchoolLevelCondition: schemeData.Criteria.HasChildren.SchoolLevelCondition,
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	benefits := make([]models.Benefit, len(schemeData.Benefits))
	for i, benefit := range schemeData.Benefits {
		if err := utils.ValidateBenefit(benefit.Name, benefit.Amount); err != nil {
			tx.Rollback()
			return err
		}

		benefits[i] = models.Benefit{
			ID:       utils.GenerateUUID(),
			Name:     benefit.Name,
			Amount:   benefit.Amount,
			SchemeID: scheme.ID,
		}
	}

	if err := tx.Create(&scheme).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(benefits) > 0 {
		if err := tx.Create(&benefits).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// RETRIEVE All Schemes
func (s *SchemeService) GetAllSchemes() ([]dto.Scheme, error) {
	var schemes []models.Scheme
	if err := s.DB.Preload("Benefits").Find(&schemes).Error; err != nil {
		return nil, err
	}

	output := make([]dto.Scheme, len(schemes))
	for i, scheme := range schemes {
		benefitDTO := make([]dto.Benefit, len(scheme.Benefits))
		for j, benefit := range scheme.Benefits {
			benefitDTO[j] = dto.Benefit{
				ID:     benefit.ID,
				Name:   benefit.Name,
				Amount: benefit.Amount,
			}
		}

		output[i] = dto.Scheme{
			ID:       scheme.ID,
			Name:     scheme.Name,
			Criteria: dto.CriteriaFromModel(scheme.Criteria),
			Benefits: benefitDTO,
		}
	}

	return output, nil
}

// RETRIEVE Scheme by ID
func (s *SchemeService) GetSchemeByID(id string) (*dto.Scheme, error) {
	var scheme models.Scheme
	if err := s.DB.Preload("Benefits").First(&scheme, "id = ?", id).Error; err != nil {
		return nil, errors.New("scheme not found")
	}

	benefitDTO := make([]dto.Benefit, len(scheme.Benefits))
	for i, benefit := range scheme.Benefits {
		benefitDTO[i] = dto.Benefit{
			ID:     benefit.ID,
			Name:   benefit.Name,
			Amount: benefit.Amount,
		}
	}

	return &dto.Scheme{
		ID:       scheme.ID,
		Name:     scheme.Name,
		Criteria: dto.CriteriaFromModel(scheme.Criteria),
		Benefits: benefitDTO,
	}, nil
}

// UDPATE Scheme by ID
func (s *SchemeService) UpdateScheme(id string, updatedData *models.Scheme) error {
	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := utils.ValidateScheme(
		updatedData.Name,
		updatedData.Criteria.EmploymentStatus,
		updatedData.Criteria.HasChildren,
	); err != nil {
		return err
	}

	var scheme models.Scheme

	if err := tx.First(&scheme, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return errors.New("scheme not found")
	}

	scheme.Name = updatedData.Name
	scheme.Criteria = models.Criteria{
		EmploymentStatus: updatedData.Criteria.EmploymentStatus,
		HasChildren: &models.Children{
			SchoolLevel:          updatedData.Criteria.HasChildren.SchoolLevel,
			SchoolLevelCondition: updatedData.Criteria.HasChildren.SchoolLevelCondition,
		},
	}
	scheme.UpdatedAt = time.Now()

	var updatedBenefits []models.Benefit
	for _, benefit := range updatedData.Benefits {

		if err := utils.ValidateBenefit(benefit.Name, benefit.Amount); err != nil {
			tx.Rollback()
			return err
		}
		if benefit.ID == "" {
			updatedBenefits = append(updatedBenefits, models.Benefit{
				ID:       utils.GenerateUUID(),
				Name:     benefit.Name,
				Amount:   benefit.Amount,
				SchemeID: scheme.ID,
			})
		} else {
			updatedBenefits = append(updatedBenefits, models.Benefit{
				ID:       benefit.ID,
				Name:     benefit.Name,
				Amount:   benefit.Amount,
				SchemeID: scheme.ID,
			})
		}
	}

	if err := tx.Save(&scheme).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("scheme_id = ?", scheme.ID).Delete(&models.Benefit{}).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete old benefits")
	}

	if len(updatedBenefits) > 0 {
		if err := tx.Create(&updatedBenefits).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// DELETE Scheme
func (s *SchemeService) DeleteScheme(id string) error {
	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var scheme models.Scheme

	if err := tx.First(&scheme, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return errors.New("scheme not found")
	}

	if err := tx.Where("scheme_id = ?", id).Delete(&models.Benefit{}).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete scheme benefits")
	}

	if err := tx.Delete(&scheme).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete scheme")
	}

	return tx.Commit().Error
}

// RETRIEVE Eligible Schemes
func (s *SchemeService) GetEligibleSchemes(applicantID string) ([]dto.Scheme, error) {
	var applicant models.Applicant
	if err := s.DB.Preload("Household").First(&applicant, "id = ?", applicantID).Error; err != nil {
		return nil, errors.New("applicant not found")
	}

	var schemes []models.Scheme
	if err := s.DB.Preload("Benefits").Find(&schemes).Error; err != nil {
		return nil, errors.New("failed to retrieve schemes")
	}

	eligibleSchemes := []dto.Scheme{}
	for _, scheme := range schemes {
		if isEligible(applicant, scheme.Criteria) {
			benefitDTO := make([]dto.Benefit, len(scheme.Benefits))
			for i, benefit := range scheme.Benefits {
				benefitDTO[i] = dto.Benefit{
					ID:     benefit.ID,
					Name:   benefit.Name,
					Amount: benefit.Amount,
				}
			}

			eligibleSchemes = append(eligibleSchemes, dto.Scheme{
				ID:       scheme.ID,
				Name:     scheme.Name,
				Criteria: dto.CriteriaFromModel(scheme.Criteria),
				Benefits: benefitDTO,
			})
		}
	}

	return eligibleSchemes, nil
}
