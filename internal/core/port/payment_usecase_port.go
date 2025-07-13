package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
)

type PaymentUseCase interface {
	Create(ctx context.Context, input dto.CreatePaymentInput) (*entity.Payment, error)
	Update(ctx context.Context, payment dto.UpdatePaymentInput) (*entity.Payment, error)
	Get(ctx context.Context, payment dto.GetPaymentInput) (*entity.Payment, error)
}
