package models

import "time"

type Application struct {
	ID          string    `json:"id" gorm:"type:uuid;primaryKey"`
	ApplicantID string    `json:"applicant_id" gorm:"type:uuid;not null;index"`
	SchemeID    string    `json:"scheme_id" gorm:"type:uuid;not null;index"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
