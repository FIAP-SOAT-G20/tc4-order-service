package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
)

type CustomerController interface {
	List(ctx context.Context, presenter Presenter, input dto.ListCustomersInput) ([]byte, error)
	Create(ctx context.Context, presenter Presenter, input dto.CreateCustomerInput) ([]byte, error)
	Get(ctx context.Context, presenter Presenter, input dto.GetCustomerInput) ([]byte, error)
	Update(ctx context.Context, presenter Presenter, input dto.UpdateCustomerInput) ([]byte, error)
	Delete(ctx context.Context, presenter Presenter, input dto.DeleteCustomerInput) ([]byte, error)
}
