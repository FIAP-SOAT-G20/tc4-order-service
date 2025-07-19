package request

import valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"

type ListStaffsQueryRequest struct {
	Name  string                `form:"name" example:"John Doe"`
	Role  valueobject.StaffRole `form:"role" binding:"omitempty" example:"COOK"`
	Page  int                   `form:"page,default=1" example:"1"`
	Limit int                   `form:"limit,default=10" example:"10"`
}

type CreateStaffBodyRequest struct {
	Name string                `json:"name" binding:"required,min=3,max=100" example:"John Doe"`
	Role valueobject.StaffRole `json:"role" binding:"required,staff_role_exists,max=500" example:"COOK"`
}

type GetStaffUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UpdateStaffUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UpdateStaffBodyRequest struct {
	Name string                `json:"name" binding:"required,min=3,max=100" example:"John Doe"`
	Role valueobject.StaffRole `json:"role" binding:"required,staff_role_exists,max=500" example:"COOK"`
}

type DeleteStaffUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}
