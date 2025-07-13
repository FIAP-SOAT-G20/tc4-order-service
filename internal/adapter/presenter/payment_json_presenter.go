package presenter

import (
	"encoding/json"
	"errors"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type paymentJsonPresenter struct{}

// PaymentJsonResponse represents the response of a payment
func NewPaymentJsonPresenter() port.Presenter {
	return &paymentJsonPresenter{}
}

// Present write the response to the client
func (p *paymentJsonPresenter) Present(pp dto.PresenterInput) ([]byte, error) {
	switch v := pp.Result.(type) {
	case *entity.Payment:
		output := ToPaymentJsonResponse(v)
		return json.Marshal(output)
	case []*entity.Payment:
		paymentOutputs := make([]PaymentJsonResponse, len(v))
		for i, payment := range v {
			paymentOutputs[i] = ToPaymentJsonResponse(payment)
		}

		output := &PaymentJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Payments: paymentOutputs,
		}

		return json.Marshal(output)
	default:
		return nil, domain.NewInternalError(errors.New(domain.ErrInternalError))
	}
}

// ToPaymentJsonResponse convert entity.Payment to PaymentJsonResponse
func ToPaymentJsonResponse(p *entity.Payment) PaymentJsonResponse {
	return PaymentJsonResponse{
		ID:                p.ID,
		Status:            p.Status,
		OrderID:           p.OrderID,
		ExternalPaymentID: p.ExternalPaymentID,
		QrData:            p.QrData,
	}
}
