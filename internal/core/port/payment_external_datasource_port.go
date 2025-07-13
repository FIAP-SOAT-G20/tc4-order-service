package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
)

type PaymentExternalDatasource interface {
	Create(context context.Context, payment *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error)
}
