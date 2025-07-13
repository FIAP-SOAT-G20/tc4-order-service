package entity

import (
	"time"

	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
)

type Payment struct {
	ID                uint64
	Status            valueobject.PaymentStatus
	ExternalPaymentID string
	QrData            string
	OrderID           uint64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type CreatePaymentExternalInput struct {
	ExternalReference string
	TotalAmount       float32
	Items             []PaymentExternalItemsInput
	Title             string
	Description       string
	NotificationUrl   string
}

type PaymentExternalItemsInput struct {
	Category    string
	Title       string
	Description string
	UnitPrice   float32
	Quantity    uint64
	UnitMeasure string
	TotalAmount float32
}

type CreatePaymentExternalOutput struct {
	InStoreOrderID string
	QrData         string
}
