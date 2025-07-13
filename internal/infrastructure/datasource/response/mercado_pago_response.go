package datasource_response

import "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"

type MercadoPagoResponse struct {
	InStoreOrderID string `json:"in_store_order_id"`
	QrData         string `json:"qr_data"`
}

func NewCreatePaymentExternalOutput(r *MercadoPagoResponse) *entity.CreatePaymentExternalOutput {
	return &entity.CreatePaymentExternalOutput{
		InStoreOrderID: r.InStoreOrderID,
		QrData:         r.QrData,
	}
}
