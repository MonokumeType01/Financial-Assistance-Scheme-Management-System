package models

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateUUID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

type Applicant struct {
	ID               string `json:"id" gorm:"type:char(32);primaryKey"`
	Name             string `json:"name"`
	EmploymentStatus string `json:"employment_status"`
	Sex              string `json:"sex"`
	DateOfBirth      string `json:"date_of_birth"`
}

type ApplicantWithHouseHold struct {
	Applicant
	Household []HouseholdMember `json:"household" gorm:"foreignKey:ApplicantID"`
}

type HouseholdMember struct {
	Applicant
	Relation    string `json:"relation"`
	ApplicantID string `json:"applicant_id"`
}
