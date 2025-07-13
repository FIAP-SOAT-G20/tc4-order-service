package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type authController struct {
	authUseCase port.AuthUseCase
}

func NewAuthController(authUseCase port.AuthUseCase) port.AuthController {
	return &authController{
		authUseCase: authUseCase,
	}
}

func (c *authController) Authenticate(ctx context.Context, presenter port.Presenter, input dto.AuthenticateInput) ([]byte, error) {
	token, err := c.authUseCase.Authenticate(ctx, input)
	if err != nil {
		return nil, err
	}

	return presenter.Present(dto.PresenterInput{Result: token})
}
