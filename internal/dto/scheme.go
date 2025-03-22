package dto

import (
	"fmt"

	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/data"
	"github.com/MonokumeType01/Financial-Assistance-Scheme-Management-System/internal/models"
)

func ChildrenFromModel(children *models.Children) *Children {
	if children == nil {
		return nil
	}

	schoolLevelDisplay := fmt.Sprintf(
		"%s %s",
		data.CRITERIA_TYPE_ID_MAP[children.SchoolLevelCondition],
		data.SCHOOL_LEVEL_TYPE_ID_MAP[children.SchoolLevel],
	)

	return &Children{
		SchoolLevel: schoolLevelDisplay,
	}
}

func ChildrenToModel(children *Children) *models.Children {
	if children == nil {
		return nil
	}

	var conditionStr, levelStr string
	fmt.Sscanf(children.SchoolLevel, "%s %s", &conditionStr, &levelStr)

	// Reverse Mapping for Database Storage
	schoolLevel := 0
	for key, value := range data.SCHOOL_LEVEL_TYPE_ID_MAP {
		if value == levelStr {
			schoolLevel = key
			break
		}
	}

	conditionVal := 0
	for key, value := range data.CRITERIA_TYPE_ID_MAP {
		if value == conditionStr {
			conditionVal = key
			break
		}
	}

	return &models.Children{
		SchoolLevel:          schoolLevel,
		SchoolLevelCondition: conditionVal,
	}
}

func CriteriaFromModel(criteria models.Criteria) Criteria {
	return Criteria{
		EmploymentStatus: criteria.EmploymentStatus,
		HasChildren:      ChildrenFromModel(criteria.HasChildren),
	}
}

type Benefit struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type Scheme struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Criteria Criteria  `json:"criteria,omitempty"`
	Benefits []Benefit `json:"benefits"`
}

type Criteria struct {
	EmploymentStatus string    `json:"employment_status"`
	HasChildren      *Children `json:"has_children,omitempty"`
}

type Children struct {
	SchoolLevel          string `json:"school_level"`
	SchoolLevelCondition string `json:"-"`
}
