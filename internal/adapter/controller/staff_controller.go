package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type StaffController struct {
	useCase port.StaffUseCase
}

func NewStaffController(useCase port.StaffUseCase) port.StaffController {
	return &StaffController{useCase}
}

func (c *StaffController) List(ctx context.Context, p port.Presenter, i dto.ListStaffsInput) ([]byte, error) {
	staffs, total, err := c.useCase.List(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{
		Total:  total,
		Page:   i.Page,
		Limit:  i.Limit,
		Result: staffs,
	})
}

func (c *StaffController) Create(ctx context.Context, p port.Presenter, i dto.CreateStaffInput) ([]byte, error) {
	staff, err := c.useCase.Create(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: staff})
}

func (c *StaffController) Get(ctx context.Context, p port.Presenter, i dto.GetStaffInput) ([]byte, error) {
	staff, err := c.useCase.Get(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: staff})
}

func (c *StaffController) Update(ctx context.Context, p port.Presenter, i dto.UpdateStaffInput) ([]byte, error) {
	staff, err := c.useCase.Update(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: staff})
}

func (c *StaffController) Delete(ctx context.Context, p port.Presenter, i dto.DeleteStaffInput) ([]byte, error) {
	staff, err := c.useCase.Delete(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: staff})
}
