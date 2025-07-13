package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
)

type PaymentDataSource interface {
	Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
	GetByOrderID(ctx context.Context, orderID uint64) (*entity.Payment, error)
	GetByOrderIDAndStatusProcessing(ctx context.Context, orderID uint64) (*entity.Payment, error)
	UpdateStatus(ctx context.Context, status valueobject.PaymentStatus, externalPaymentID string) error
	GetByExternalPaymentID(ctx context.Context, externalPaymentID string) (*entity.Payment, error)
}
