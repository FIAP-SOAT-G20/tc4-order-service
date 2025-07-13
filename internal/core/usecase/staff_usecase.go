package usecase

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type staffUseCase struct {
	gateway port.StaffGateway
}

// NewStaffUseCase creates a new StaffUseCase
func NewStaffUseCase(gateway port.StaffGateway) port.StaffUseCase {
	return &staffUseCase{gateway: gateway}
}

// List returns a list of staffs
func (uc *staffUseCase) List(ctx context.Context, i dto.ListStaffsInput) ([]*entity.Staff, int64, error) {
	staffs, total, err := uc.gateway.FindAll(ctx, i.Name, i.Role, i.Page, i.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return staffs, total, nil
}

// Create creates a new staff
func (uc *staffUseCase) Create(ctx context.Context, i dto.CreateStaffInput) (*entity.Staff, error) {
	staff := i.ToEntity()

	if err := uc.gateway.Create(ctx, staff); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}

// Get returns a staff by ID
func (uc *staffUseCase) Get(ctx context.Context, i dto.GetStaffInput) (*entity.Staff, error) {
	staff, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return staff, nil
}

// Update updates a staff
func (uc *staffUseCase) Update(ctx context.Context, i dto.UpdateStaffInput) (*entity.Staff, error) {
	staff, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	staff.Update(i.Name, i.Role)

	if err := uc.gateway.Update(ctx, staff); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}

// Delete deletes a staff
func (uc *staffUseCase) Delete(ctx context.Context, i dto.DeleteStaffInput) (*entity.Staff, error) {
	staff, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, i.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}
