package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type paymentGayeway struct {
	dataSource       port.PaymentDataSource
	dataSourceRemote port.PaymentExternalDatasource
}

func NewPaymentGateway(
	dataSource port.PaymentDataSource,
	dataSourceRemote port.PaymentExternalDatasource,
) port.PaymentGateway {
	return &paymentGayeway{dataSource, dataSourceRemote}
}

func (g *paymentGayeway) FindByOrderID(ctx context.Context, orderID uint64) (*entity.Payment, error) {
	return g.dataSource.GetByOrderID(ctx, orderID)
}

func (g *paymentGayeway) FindByOrderIDAndStatusProcessing(ctx context.Context, orderID uint64) (*entity.Payment, error) {
	return g.dataSource.GetByOrderIDAndStatusProcessing(ctx, orderID)
}

func (g *paymentGayeway) Create(ctx context.Context, p *entity.Payment) (*entity.Payment, error) {
	return g.dataSource.Create(ctx, p)
}

func (g *paymentGayeway) FindByExternalPaymentID(ctx context.Context, externalPaymentId string) (*entity.Payment, error) {
	return g.dataSource.GetByExternalPaymentID(ctx, externalPaymentId)
}

func (g *paymentGayeway) Update(ctx context.Context, status valueobject.PaymentStatus, resource string) error {
	return g.dataSource.UpdateStatus(ctx, status, resource)
}

func (g *paymentGayeway) CreateExternal(ctx context.Context, payment *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error) {
	return g.dataSourceRemote.Create(ctx, payment)
}
