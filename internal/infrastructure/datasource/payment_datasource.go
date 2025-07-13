package datasource

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type paymentDataSource struct {
	db *gorm.DB
}

type PaymentKey string

func NewPaymentDataSource(db *gorm.DB) port.PaymentDataSource {
	return &paymentDataSource{db}
}

func (ds *paymentDataSource) Create(ctx context.Context, p *entity.Payment) (*entity.Payment, error) {
	if err := ds.db.WithContext(ctx).Create(p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (ds *paymentDataSource) GetByOrderID(ctx context.Context, orderID uint64) (*entity.Payment, error) {
	var p entity.Payment
	if err := ds.db.WithContext(ctx).Where("order_id = ?", orderID).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (ds *paymentDataSource) GetByOrderIDAndStatusProcessing(ctx context.Context, orderID uint64) (*entity.Payment, error) {
	var p entity.Payment
	if err := ds.db.WithContext(ctx).Where("order_id = ? AND status = ?", orderID, valueobject.PROCESSING).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &p, nil
		}
		return nil, err
	}

	return &p, nil
}

func (ds *paymentDataSource) UpdateStatus(ctx context.Context, status valueobject.PaymentStatus, epID string) error {
	if err := ds.db.WithContext(ctx).Model(&entity.Payment{}).Where("external_payment_id = ?", epID).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}

func (ds *paymentDataSource) GetByExternalPaymentID(ctx context.Context, epID string) (*entity.Payment, error) {
	var payment entity.Payment

	if err := ds.db.WithContext(ctx).Where("external_payment_id = ?", epID).First(&payment); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	return &payment, nil
}
