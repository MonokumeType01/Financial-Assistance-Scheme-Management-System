package models

import "time"

type Applicant struct {
	ID               string            `json:"id" gorm:"type:uuid;primaryKey"`
	Name             string            `json:"name"`
	EmploymentStatus string            `json:"employment_status"`
	Sex              string            `json:"sex"`
	DateOfBirth      string            `json:"date_of_birth"`
	Household        []HouseholdMember `gorm:"foreignKey:ApplicantID;constraint:OnDelete:CASCADE"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}

type HouseholdMember struct {
	ID               string `json:"id" gorm:"type:uuid;primaryKey"`
	Name             string `json:"name"`
	EmploymentStatus string `json:"employment_status"`
	Sex              string `json:"sex"`
	DateOfBirth      string `json:"date_of_birth"`
	Relation         string `json:"relation"`
	ApplicantID      string `json:"-" gorm:"type:uuid;index;not null"`
	SchoolLevel      int    `json:"school_level" gorm:"type:int"`
}

type ApplicantWithHousehold struct {
	Applicant
	Household []HouseholdMember `json:"household"`
}
