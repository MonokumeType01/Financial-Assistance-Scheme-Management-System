package dto

type HouseholdMember struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	EmploymentStatus string `json:"employment_status"`
	Sex              string `json:"sex"`
	DateOfBirth      string `json:"date_of_birth"`
	Relation         string `json:"relation"`
	SchoolLevel      int    `json:"-"`
}

type Applicant struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	EmploymentStatus string `json:"employment_status"`
	Sex              string `json:"sex"`
	DateOfBirth      string `json:"date_of_birth"`
}

type ApplicantWithHousehold struct {
	Applicant
	Household []HouseholdMember `json:"household"`
}
