package dto

import (
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
)

type CreateStaffInput struct {
	Name string
	Role valueobject.StaffRole
}

func (i CreateStaffInput) ToEntity() *entity.Staff {
	return &entity.Staff{
		Name: i.Name,
		Role: i.Role,
	}
}

type UpdateStaffInput struct {
	ID   uint64
	Name string
	Role valueobject.StaffRole
}

type GetStaffInput struct {
	ID uint64
}

type DeleteStaffInput struct {
	ID uint64
}

type ListStaffsInput struct {
	Name  string
	Role  valueobject.StaffRole
	Page  int
	Limit int
}
