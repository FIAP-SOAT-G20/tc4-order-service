package usecase

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type authUseCase struct {
	customerUseCase port.CustomerUseCase
	jwtService      port.JWTService
}

// NewAuthUseCase creates a new auth use case instance
func NewAuthUseCase(customerUseCase port.CustomerUseCase, jwtService port.JWTService) port.AuthUseCase {
	return &authUseCase{
		customerUseCase: customerUseCase,
		jwtService:      jwtService,
	}
}

// Authenticate authenticates a customer by CPF and returns the customer entity, token and expiration time
func (u *authUseCase) Authenticate(ctx context.Context, input dto.AuthenticateInput) (string, error) {
	// Find customer by CPF
	customer, err := u.customerUseCase.FindByCPF(ctx, dto.FindCustomerByCPFInput(input))
	if err != nil {
		return "", err
	}

	// Generate JWT token - let the service handle the token duration
	token, err := u.jwtService.GenerateToken(customer.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
