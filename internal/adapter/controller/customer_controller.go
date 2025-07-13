package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type customerController struct {
	useCase port.CustomerUseCase
}

func NewCustomerController(useCase port.CustomerUseCase) port.CustomerController {
	return &customerController{useCase}
}

func (c *customerController) List(ctx context.Context, p port.Presenter, i dto.ListCustomersInput) ([]byte, error) {
	customers, total, err := c.useCase.List(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{
		Total:  total,
		Page:   i.Page,
		Limit:  i.Limit,
		Result: customers,
	})
}

func (c *customerController) Create(ctx context.Context, p port.Presenter, i dto.CreateCustomerInput) ([]byte, error) {
	customer, err := c.useCase.Create(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: customer})
}

func (c *customerController) Get(ctx context.Context, p port.Presenter, i dto.GetCustomerInput) ([]byte, error) {
	customer, err := c.useCase.Get(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: customer})
}

func (c *customerController) Update(ctx context.Context, p port.Presenter, i dto.UpdateCustomerInput) ([]byte, error) {
	customer, err := c.useCase.Update(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: customer})
}

func (c *customerController) Delete(ctx context.Context, p port.Presenter, i dto.DeleteCustomerInput) ([]byte, error) {
	customer, err := c.useCase.Delete(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: customer})
}
