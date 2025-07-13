package entity

import (
	"time"

	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
)

type Staff struct {
	ID        uint64
	Name      string
	Role      valueobject.StaffRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStaff(name string, role valueobject.StaffRole) *Staff {
	staff := &Staff{
		Name:      name,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return staff
}

func (p *Staff) Update(name string, role valueobject.StaffRole) {
	p.Name = name
	p.Role = role
	p.UpdatedAt = time.Now()
}
