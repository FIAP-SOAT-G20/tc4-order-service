package response

import "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"

type CreatePaymentResponse struct {
	InStoreOrderID string `json:"in_store_order_id"`
	QrData         string `json:"qr_data"`
}

func (r *CreatePaymentResponse) ToEntity() *entity.CreatePaymentExternalOutput {
	return &entity.CreatePaymentExternalOutput{
		InStoreOrderID: r.InStoreOrderID,
		QrData:         r.QrData,
	}
}
