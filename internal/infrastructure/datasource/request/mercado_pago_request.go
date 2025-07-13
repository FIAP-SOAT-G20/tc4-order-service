package datasource_request

import "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"

type FakeMercadoPagoRequest struct {
	ExternalReference string                        `json:"external_reference"`
	NotificationUrl   string                        `json:"notification_url"`
	TotalAmount       float32                       `json:"total_amount"`
	Title             string                        `json:"title"`
	Description       string                        `json:"description"`
	Items             []FakeMercadoPagoItemsRequest `json:"items"`
}

type FakeMercadoPagoItemsRequest struct {
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    uint64  `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float32 `json:"total_amount"`
}

func NewFakeMercadoPagoRequest(p *entity.CreatePaymentExternalInput) *FakeMercadoPagoRequest {
	fake := &FakeMercadoPagoRequest{
		ExternalReference: p.ExternalReference,
		TotalAmount:       p.TotalAmount,
		Title:             p.Title,
		Description:       p.Description,
		NotificationUrl:   p.NotificationUrl,
	}
	for _, item := range p.Items {
		fake.Items = append(fake.Items, *newFakeMercadoPagoItemsRequest(&item))
	}
	return fake
}

func newFakeMercadoPagoItemsRequest(p *entity.PaymentExternalItemsInput) *FakeMercadoPagoItemsRequest {
	return &FakeMercadoPagoItemsRequest{
		Title:       p.Title,
		Description: p.Description,
		UnitPrice:   p.UnitPrice,
		Category:    p.Category,
		UnitMeasure: p.UnitMeasure,
		Quantity:    p.Quantity,
		TotalAmount: p.TotalAmount,
	}
}
