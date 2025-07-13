package valueobject

import "strings"

type StaffRole string

const (
	COOK      StaffRole = "COOK"
	ATTENDANT StaffRole = "ATTENDANT"
	MANAGER   StaffRole = "MANAGER"
	UNDEFINED StaffRole = ""
)

func IsValidStaffRole(status string) bool {
	return ToStaffRole(status) != UNDEFINED
}

func (o StaffRole) String() string {
	return strings.ToUpper(string(o))
}

func ToStaffRole(status string) StaffRole {
	switch strings.ToUpper(status) {
	case "COOK":
		return COOK
	case "ATTENDANT":
		return ATTENDANT
	case "MANAGER":
		return MANAGER
	default:
		return UNDEFINED
	}
}
