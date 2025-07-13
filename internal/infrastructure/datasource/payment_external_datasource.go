package datasource

import (
	"context"
	"fmt"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/infrastructure/handler/request"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/infrastructure/handler/response"
	"github.com/go-resty/resty/v2"
)

type PaymentExternalDataSource struct {
	httpClient *resty.Client
}

func NewPaymentExternalDataSource(client *resty.Client) port.PaymentExternalDatasource {
	return &PaymentExternalDataSource{client}
}

func (ds *PaymentExternalDataSource) Create(ctx context.Context, p *entity.CreatePaymentExternalInput) (*entity.CreatePaymentExternalOutput, error) {
	cfg := config.LoadConfig()
	p.NotificationUrl = cfg.MercadoPagoNotificationURL
	paymentRequest := request.NewPaymentRequest(p)

	var result response.CreatePaymentResponse
	resp, err := ds.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+cfg.MercadoPagoToken).
		SetBody(paymentRequest).
		SetResult(&result).
		Post(cfg.MercadoPagoURL)
	if err != nil {
		return nil, fmt.Errorf("error to create payment: %w", err)
	}

	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("error: response status %d", resp.StatusCode())
	}

	return result.ToEntity(), nil
}
