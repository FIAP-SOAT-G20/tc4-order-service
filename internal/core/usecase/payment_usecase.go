package usecase

import (
	"context"
	"strconv"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type paymentUseCase struct {
	paymentGateway port.PaymentGateway
	orderUseCase   port.OrderUseCase
}

// NewPaymentUseCase create a new payment use case
func NewPaymentUseCase(
	paymentGateway port.PaymentGateway,
	orderUseCase port.OrderUseCase,
) port.PaymentUseCase {
	return &paymentUseCase{paymentGateway, orderUseCase}
}

// Create create a new payment
func (uc *paymentUseCase) Create(ctx context.Context, i dto.CreatePaymentInput) (*entity.Payment, error) {
	existentPedingPayment, err := uc.paymentGateway.FindByOrderIDAndStatusProcessing(ctx, i.OrderID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if existentPedingPayment.ID != 0 {
		return existentPedingPayment, nil
	}

	order, err := uc.orderUseCase.Get(ctx, dto.GetOrderInput{ID: i.OrderID})
	if err != nil {
		return nil, domain.NewNotFoundError(domain.ErrOrderIsMandatory)
	}

	if len(order.OrderProducts) == 0 {
		return nil, domain.NewNotFoundError(domain.ErrOrderWithoutProducts)
	}

	paymentPayload := uc.createPaymentPayload(order)

	extPayment, err := uc.paymentGateway.CreateExternal(ctx, paymentPayload)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	iPayment := &entity.Payment{
		ExternalPaymentID: extPayment.InStoreOrderID,
		OrderID:           i.OrderID,
		QrData:            extPayment.QrData,
		Status:            valueobject.PROCESSING,
	}

	payment, err := uc.paymentGateway.Create(ctx, iPayment)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	orderInput := dto.UpdateOrderInput{
		ID:         order.ID,
		Status:     valueobject.PENDING,
		CustomerID: order.CustomerID,
	}

	if _, err := uc.orderUseCase.Update(ctx, orderInput); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return payment, nil
}

func (uc *paymentUseCase) Update(ctx context.Context, p dto.UpdatePaymentInput) (*entity.Payment, error) {
	if err := uc.paymentGateway.Update(ctx, valueobject.CONFIRMED, p.Resource); err != nil {
		return nil, err
	}

	paymentOUT, err := uc.paymentGateway.FindByExternalPaymentID(ctx, p.Resource)
	if err != nil {
		return nil, err
	}

	order, err := uc.orderUseCase.Get(ctx, dto.GetOrderInput{ID: paymentOUT.OrderID})
	if err != nil {
		return nil, err
	}

	orderInput := dto.UpdateOrderInput{
		ID:         order.ID,
		Status:     valueobject.RECEIVED,
		CustomerID: order.CustomerID,
	}

	if _, err := uc.orderUseCase.Update(ctx, orderInput); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return paymentOUT, nil
}

func (uc *paymentUseCase) createPaymentPayload(o *entity.Order) *entity.CreatePaymentExternalInput {
	var totalAmount float32
	var items []entity.PaymentExternalItemsInput
	externalReference := strconv.FormatUint(o.ID, 10)

	for _, v := range o.OrderProducts {
		items = append(items, entity.PaymentExternalItemsInput{
			Title:       v.Product.Name,
			Description: v.Product.Description,
			UnitPrice:   float32(v.Product.Price),
			Category:    "marketplace",
			UnitMeasure: "unit",
			Quantity:    uint64(v.Quantity),
			TotalAmount: float32(v.Product.Price) * float32(v.Quantity),
		})
		totalAmount += float32(v.Product.Price) * float32(v.Quantity)
	}

	return &entity.CreatePaymentExternalInput{
		ExternalReference: externalReference,
		TotalAmount:       totalAmount,
		Items:             items,
		Title:             "FIAP Tech Challenge - Product Order",
		Description:       "Purchases made at the FIAP Tech Challenge store",
	}
}

func (uc *paymentUseCase) Get(ctx context.Context, input dto.GetPaymentInput) (*entity.Payment, error) {
	payment, err := uc.paymentGateway.FindByOrderID(ctx, input.OrderID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if payment == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return payment, nil
}
