package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type PaymentController struct {
	useCase port.PaymentUseCase
}

func NewPaymentController(useCase port.PaymentUseCase) port.PaymentController {
	return &PaymentController{useCase}
}

func (c *PaymentController) Create(ctx context.Context, p port.Presenter, i dto.CreatePaymentInput) ([]byte, error) {
	payment, err := c.useCase.Create(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: payment})
}

func (c *PaymentController) Update(ctx context.Context, p port.Presenter, i dto.UpdatePaymentInput) ([]byte, error) {
	payment, err := c.useCase.Update(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: payment})
}

func (c *PaymentController) Get(ctx context.Context, p port.Presenter, i dto.GetPaymentInput) ([]byte, error) {
	payment, err := c.useCase.Get(ctx, i)
	if err != nil {
		return nil, err
	}

	return p.Present(dto.PresenterInput{Result: payment})
}
