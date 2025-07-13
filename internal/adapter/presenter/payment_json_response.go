package presenter

import valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"

type PaymentJsonResponse struct {
	ID                uint64                    `json:"id" example:"1"`
	Status            valueobject.PaymentStatus `json:"status" example:"pending"`
	OrderID           uint64                    `json:"order_id" example:"1"`
	ExternalPaymentID string                    `json:"external_payment_id" example:"a0aa0f26-6e0a-4b90-8c49-9f1a9c03ebcc"`
	QrData            string                    `json:"qr_data" example:"qr_data_a0aa0f26-6e0a-4b90-8c49-9f1a9c03ebcc"`
}

type PaymentJsonPaginatedResponse struct {
	JsonPagination
	Payments []PaymentJsonResponse `json:"payments"`
}
