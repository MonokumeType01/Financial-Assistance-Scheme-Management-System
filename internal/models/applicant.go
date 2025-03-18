package models

type Applicant struct {
	ID               string `json:"id" gorm:"type:uuid;primaryKey"`
	Name             string `json:"name"`
	EmploymentStatus string `json:"employment_status"`
	Sex              string `json:"sex"`
	DateOfBirth      string `json:"date_of_birth"`
}

type HouseholdMember struct {
	ID               string `json:"id" gorm:"type:uuid;primaryKey"`
	Name             string `json:"name"`
	EmploymentStatus string `json:"employment_status"`
	Sex              string `json:"sex"`
	DateOfBirth      string `json:"date_of_birth"`
	Relation         string `json:"relation"`
	ApplicantID      string `json:"-" gorm:"type:uuid;index;not null"`
}

type ApplicantWithHouseHold struct {
	Applicant
	Household []HouseholdMember `json:"household" gorm:"foreignKey:ApplicantID;constraint:OnDelete:RESTRICT"`
}
