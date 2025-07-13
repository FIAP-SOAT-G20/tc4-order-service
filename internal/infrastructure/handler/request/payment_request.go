package request

import "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"

type CreatePaymentUriRequest struct {
	OrderID uint64 `uri:"order_id" binding:"required"`
}

type CreatePaymentRequest struct {
	ExternalReference string         `json:"external_reference"`
	TotalAmount       float32        `json:"total_amount"`
	Items             []ItemsRequest `json:"items"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	NotificationURL   string         `json:"notification_url"`
}

type ItemsRequest struct {
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    uint64  `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float32 `json:"total_amount"`
}

func NewPaymentRequest(payment *entity.CreatePaymentExternalInput) *CreatePaymentRequest {
	if payment == nil {
		return nil
	}

	items := make([]ItemsRequest, 0)
	for _, item := range payment.Items {
		items = append(items, ItemsRequest{
			Title:       item.Title,
			Description: item.Description,
			UnitPrice:   item.UnitPrice,
			Category:    item.Category,
			UnitMeasure: item.UnitMeasure,
			Quantity:    item.Quantity,
			TotalAmount: item.TotalAmount,
		})
	}

	return &CreatePaymentRequest{
		ExternalReference: payment.ExternalReference,
		TotalAmount:       payment.TotalAmount,
		Items:             items,
		Title:             payment.Title,
		Description:       payment.Description,
		NotificationURL:   payment.NotificationUrl,
	}
}

type UpdatePaymentRequest struct {
	Resource string `json:"resource" binding:"required"`
	Topic    string `json:"topic" binding:"required"`
}

type GetPaymentRequest struct {
	OrderID uint64 `uri:"order_id" binding:"required"`
}
