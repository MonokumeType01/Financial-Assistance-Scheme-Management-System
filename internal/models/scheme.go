package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Children struct {
	SchoolLevel          int `json:"school_level"`
	SchoolLevelCondition int `json:"school_level_condition"` // e.g. "==", "<=", ">="
}

type Criteria struct {
	EmploymentStatus string    `json:"employment_status"`
	HasChildren      *Children `json:"has_children,omitempty"`
}

type Benefit struct {
	ID       string  `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	SchemeID string  `json:"scheme_id" gorm:"type:uuid;not null"`
}

type Scheme struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string    `json:"name"`
	Criteria  Criteria  `json:"criteria" gorm:"type:jsonb"`
	Benefits  []Benefit `json:"benefits" gorm:"foreignKey:SchemeID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Criteria) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONB value")
	}
	return json.Unmarshal(bytes, &c)
}

func (c Criteria) Value() (driver.Value, error) {
	return json.Marshal(c)
}
