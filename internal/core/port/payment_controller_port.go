package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
)

type PaymentController interface {
	Create(ctx context.Context, presenter Presenter, input dto.CreatePaymentInput) ([]byte, error)
	Update(ctx context.Context, p Presenter, i dto.UpdatePaymentInput) ([]byte, error)
	Get(ctx context.Context, p Presenter, i dto.GetPaymentInput) ([]byte, error)
}
