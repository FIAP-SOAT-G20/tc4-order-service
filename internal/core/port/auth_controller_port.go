package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
)

type AuthController interface {
	Authenticate(ctx context.Context, presenter Presenter, input dto.AuthenticateInput) ([]byte, error)
}
