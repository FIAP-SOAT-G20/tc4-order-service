package usecase

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type customerUseCase struct {
	gateway port.CustomerGateway
}

// NewCustomerUseCase creates a new CreateCustomerUseCase
func NewCustomerUseCase(gateway port.CustomerGateway) port.CustomerUseCase {
	return &customerUseCase{gateway}
}

// List returns a list of Customers
func (uc *customerUseCase) List(ctx context.Context, i dto.ListCustomersInput) ([]*entity.Customer, int64, error) {
	customers, total, err := uc.gateway.FindAll(ctx, i.Name, i.Page, i.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return customers, total, nil
}

// Create creates a new Customer
func (uc *customerUseCase) Create(ctx context.Context, i dto.CreateCustomerInput) (*entity.Customer, error) {
	customer := i.ToEntity()

	if err := uc.gateway.Create(ctx, customer); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return customer, nil
}

// Get returns a Customer by ID
func (uc *customerUseCase) Get(ctx context.Context, i dto.GetCustomerInput) (*entity.Customer, error) {
	customer, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if customer == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return customer, nil
}

// Update updates a Customer
func (uc *customerUseCase) Update(ctx context.Context, i dto.UpdateCustomerInput) (*entity.Customer, error) {
	customer, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if customer == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	customer.Update(i.Name, i.Email)

	if err := uc.gateway.Update(ctx, customer); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return customer, nil
}

// Delete deletes a Customer
func (uc *customerUseCase) Delete(ctx context.Context, i dto.DeleteCustomerInput) (*entity.Customer, error) {
	customer, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if customer == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, i.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return customer, nil
}

// FindByCPF returns a Customer by CPF
func (uc *customerUseCase) FindByCPF(ctx context.Context, input dto.FindCustomerByCPFInput) (*entity.Customer, error) {
	customer, err := uc.gateway.FindByCPF(ctx, input.CPF)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if customer == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return customer, nil
}
