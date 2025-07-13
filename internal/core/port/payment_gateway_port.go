package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
)

type PaymentGateway interface {
	Create(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
	CreateExternal(ctx context.Context, payment *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error)
	FindByOrderID(ctx context.Context, orderID uint64) (*entity.Payment, error)
	FindByOrderIDAndStatusProcessing(ctx context.Context, orderID uint64) (*entity.Payment, error) // TODO: Unify with FindByExternalPaymentID into FindOne
	FindByExternalPaymentID(ctx context.Context, resource string) (*entity.Payment, error)         // TODO: Unify with FindByExternalPaymentID into FindOne
	Update(ctx context.Context, status valueobject.PaymentStatus, resource string) error
}
