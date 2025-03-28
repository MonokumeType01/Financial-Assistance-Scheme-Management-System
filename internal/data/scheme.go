package data

// School Level Constants
const (
	SCHOOL_LEVEL_PRESCHOOL = iota + 1
	SCHOOL_LEVEL_PRIMARY
	SCHOOL_LEVEL_SECONDARY
	SCHOOL_LEVEL_ITE
	SCHOOL_LEVEL_JC
	SCHOOL_LEVEL_POLYTECHNIC
	SCHOOL_LEVEL_UNIVERSITY
)

var SCHOOL_LEVEL_TYPE_ID_MAP = map[int]string{
	SCHOOL_LEVEL_PRESCHOOL:   "preschool",
	SCHOOL_LEVEL_PRIMARY:     "primary",
	SCHOOL_LEVEL_SECONDARY:   "secondary",
	SCHOOL_LEVEL_ITE:         "ite",
	SCHOOL_LEVEL_JC:          "jc",
	SCHOOL_LEVEL_POLYTECHNIC: "polytechnic",
	SCHOOL_LEVEL_UNIVERSITY:  "university",
}

// Criteria Constants
const (
	CRITERIA_EQUAL = iota + 1
	CRITERIA_EQUAL_OR_ABOVE
	CRITERIA_EQUAL_OR_BELOW
	CRITERIA_ABOVE
	CRITERIA_BELOW
)

var CRITERIA_TYPE_ID_MAP = map[int]string{
	CRITERIA_EQUAL:          "==",
	CRITERIA_EQUAL_OR_ABOVE: ">=",
	CRITERIA_EQUAL_OR_BELOW: "<=",
	CRITERIA_ABOVE:          ">",
	CRITERIA_BELOW:          "<",
}
