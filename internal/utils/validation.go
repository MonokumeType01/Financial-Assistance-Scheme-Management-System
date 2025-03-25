package utils

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/data"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
)

/* Applicant Validation */

func ValidateSchoolLevel(schoolLevel int) error {
	validLevels := map[int]bool{
		data.SCHOOL_LEVEL_PRESCHOOL:   true,
		data.SCHOOL_LEVEL_PRIMARY:     true,
		data.SCHOOL_LEVEL_SECONDARY:   true,
		data.SCHOOL_LEVEL_ITE:         true,
		data.SCHOOL_LEVEL_JC:          true,
		data.SCHOOL_LEVEL_POLYTECHNIC: true,
		data.SCHOOL_LEVEL_UNIVERSITY:  true,
	}

	if !validLevels[schoolLevel] {
		return fmt.Errorf("invalid school level provided")
	}

	return nil
}

func ValidateRelation(relation string) error {
	validRelations := map[string]bool{
		data.RELATION_SON:      true,
		data.RELATION_DAUGHTER: true,
		data.RELATION_FATHER:   true,
		data.RELATION_MOTHER:   true,
		data.RELATION_HUSBAND:  true,
		data.RELATION_WIFE:     true,
		data.RELATION_BROTHER:  true,
		data.RELATION_SISTER:   true,
	}

	if !validRelations[relation] {
		return fmt.Errorf("invalid relation: %s", relation)
	}

	return nil
}

func ValidateApplicant(name, employmentStatus, sex, dateOfBirth string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	validEmploymentStatus := map[string]bool{
		"employed":   true,
		"unemployed": true,
		"student":    true,
		"ns":         true,
	}
	if !validEmploymentStatus[employmentStatus] {
		return errors.New("invalid employment status, must be 'employed', 'student', 'ns' or 'unemployed'")
	}

	validSex := map[string]bool{
		"male":   true,
		"female": true,
	}
	if !validSex[sex] {
		return errors.New("invalid sex, must be 'male' or 'female'")
	}

	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, dateOfBirth)
	if !match {
		return errors.New("invalid date of birth format, expected YYYY-MM-DD")
	}

	dob, err := time.Parse("2006-01-02", dateOfBirth)
	if err != nil {
		return errors.New("invalid date of birth value")
	}

	if dob.After(time.Now()) {
		return errors.New("date of birth cannot be in the future")
	}

	return nil
}

/* Schemes Validation */

func ValidateScheme(name, employmentStatus string, hasChildren *models.Children) error {
	if name == "" {
		return errors.New("scheme name cannot be empty")
	}

	validEmploymentStatus := map[string]bool{
		"employed":   true,
		"unemployed": true,
	}
	if employmentStatus != "" && !validEmploymentStatus[employmentStatus] {
		return errors.New("invalid employment status, must be 'employed' or 'unemployed'")
	}

	if hasChildren != nil {
		_, isValidSchoolLevel := data.SCHOOL_LEVEL_TYPE_ID_MAP[hasChildren.SchoolLevel]
		if !isValidSchoolLevel {
			return errors.New("invalid school level provided")
		}

		_, isValidCondition := data.CRITERIA_TYPE_ID_MAP[hasChildren.SchoolLevelCondition]
		if !isValidCondition {
			return errors.New("invalid school level condition provided")
		}
	}

	return nil
}

func ValidateBenefit(name string, amount float64) error {
	if name == "" {
		return errors.New("benefit name cannot be empty")
	}

	if amount <= 0 {
		return errors.New("benefit amount must be greater than zero")
	}

	return nil
}
