package controller

import (
	"context"
	"testing"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	mockport "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPaymentController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentUseCase := mockport.NewMockPaymentUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewPaymentController(mockPaymentUseCase)

	ctx := context.Background()

	input := dto.CreatePaymentInput{
		OrderID: uint64(1),
	}

	mockPayment := &entity.Payment{}

	mockPaymentUseCase.EXPECT().
		Create(ctx, input).
		Return(mockPayment, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockPayment}).
		Return([]byte{}, nil)

	output, err := controller.Create(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestPaymentController_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentUseCase := mockport.NewMockPaymentUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewPaymentController(mockPaymentUseCase)

	ctx := context.Background()

	input := dto.GetPaymentInput{
		OrderID: uint64(1),
	}

	mockPayment := &entity.Payment{}

	mockPaymentUseCase.EXPECT().
		Get(ctx, input).
		Return(mockPayment, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockPayment}).
		Return([]byte{}, nil)

	output, err := controller.Get(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestPaymentController_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentUseCase := mockport.NewMockPaymentUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := NewPaymentController(mockPaymentUseCase)

	ctx := context.Background()

	input := dto.UpdatePaymentInput{
		Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
		Topic:    "payment",
	}

	mockPayment := &entity.Payment{}

	mockPaymentUseCase.EXPECT().
		Update(ctx, input).
		Return(mockPayment, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockPayment}).
		Return([]byte{}, nil)

	output, err := controller.Update(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}
