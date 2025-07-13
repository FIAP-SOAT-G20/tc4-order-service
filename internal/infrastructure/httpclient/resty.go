package httpclient

import (
	"fmt"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/infrastructure/config"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/infrastructure/logger"
	"github.com/go-resty/resty/v2"
)

type HTTPClient struct {
	*resty.Client
}

func NewRestyClient(cfg *config.Config, logger *logger.Logger) *HTTPClient {
	httpCLient := resty.New().
		SetTimeout(cfg.MercadoPagoTimeout).
		SetRetryCount(cfg.MercadoPagoRetryCount).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.MercadoPagoToken)).
		SetHeader("Content-Type", "application/json")

	logger.Info("resty client created")

	return &HTTPClient{httpCLient}
}
