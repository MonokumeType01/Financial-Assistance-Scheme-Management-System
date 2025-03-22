package services

import (
	"errors"

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
				if householdMember.SchoolLevel == criteria.HasChildren.SchoolLevel {
					hasEligibleChild = true
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
	conditionValue := schemeData.Criteria.HasChildren.SchoolLevelCondition

	scheme := models.Scheme{
		ID:   utils.GenerateUUID(),
		Name: schemeData.Name,
		Criteria: models.Criteria{
			EmploymentStatus: schemeData.Criteria.EmploymentStatus,
			HasChildren: &models.Children{
				SchoolLevel:          schemeData.Criteria.HasChildren.SchoolLevel,
				SchoolLevelCondition: conditionValue,
			},
		},
	}

	for _, benefit := range schemeData.Benefits {
		scheme.Benefits = append(scheme.Benefits, models.Benefit{
			ID:       utils.GenerateUUID(),
			Name:     benefit.Name,
			Amount:   benefit.Amount,
			SchemeID: scheme.ID,
		})
	}

	return s.DB.Create(&scheme).Error
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
	var scheme models.Scheme

	if err := s.DB.First(&scheme, "id = ?", id).Error; err != nil {
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

	var updatedBenefits []models.Benefit
	for _, benefit := range updatedData.Benefits {
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

	if err := s.DB.Where("scheme_id = ?", scheme.ID).Delete(&models.Benefit{}).Error; err != nil {
		return errors.New("failed to delete old benefits")
	}

	scheme.Benefits = updatedBenefits

	return s.DB.Save(&scheme).Error
}

// DELETE Scheme
func (s *SchemeService) DeleteScheme(id string) error {
	var scheme models.Scheme
	if err := s.DB.First(&scheme, "id = ?", id).Error; err != nil {
		return errors.New("scheme not found")
	}

	if err := s.DB.Where("scheme_id = ?", id).Delete(&models.Benefit{}).Error; err != nil {
		return errors.New("failed to delete scheme benefits")
	}

	return s.DB.Delete(&scheme).Error
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
