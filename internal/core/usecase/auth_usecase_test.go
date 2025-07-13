package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	mockport "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/usecase"
)

func TestAuthUseCase_Authenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		// Arrange
		mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
		mockJWTService := mockport.NewMockJWTService(ctrl)
		useCase := usecase.NewAuthUseCase(mockCustomerUseCase, mockJWTService)

		ctx := context.Background()
		input := dto.AuthenticateInput{
			CPF: "12345678901",
		}

		mockCustomer := &entity.Customer{
			ID:    1,
			Name:  "Test Customer",
			Email: "test@example.com",
			CPF:   "12345678901",
		}
		mockToken := "test-jwt-token"

		// Set up expectations
		mockCustomerUseCase.EXPECT().
			FindByCPF(ctx, dto.FindCustomerByCPFInput(input)).
			Return(mockCustomer, nil)

		mockJWTService.EXPECT().
			GenerateToken(mockCustomer.ID).
			Return(mockToken, nil)

		// Act
		token, err := useCase.Authenticate(ctx, input)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, mockToken, token)
	})

	t.Run("customer_not_found", func(t *testing.T) {
		// Arrange
		mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
		mockJWTService := mockport.NewMockJWTService(ctrl)
		useCase := usecase.NewAuthUseCase(mockCustomerUseCase, mockJWTService)

		ctx := context.Background()
		input := dto.AuthenticateInput{
			CPF: "12345678901",
		}

		expectedErr := errors.New("customer not found")

		// Set up expectations
		mockCustomerUseCase.EXPECT().
			FindByCPF(ctx, dto.FindCustomerByCPFInput(input)).
			Return(nil, expectedErr)

		// Act
		token, err := useCase.Authenticate(ctx, input)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Empty(t, token)
	})

	t.Run("token_generation_error", func(t *testing.T) {
		// Arrange
		mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
		mockJWTService := mockport.NewMockJWTService(ctrl)
		useCase := usecase.NewAuthUseCase(mockCustomerUseCase, mockJWTService)

		ctx := context.Background()
		input := dto.AuthenticateInput{
			CPF: "12345678901",
		}

		mockCustomer := &entity.Customer{
			ID:    1,
			Name:  "Test Customer",
			Email: "test@example.com",
			CPF:   "12345678901",
		}
		expectedErr := errors.New("token generation error")

		// Set up expectations
		mockCustomerUseCase.EXPECT().
			FindByCPF(ctx, dto.FindCustomerByCPFInput(input)).
			Return(mockCustomer, nil)

		mockJWTService.EXPECT().
			GenerateToken(mockCustomer.ID).
			Return("", expectedErr)

		// Act
		token, err := useCase.Authenticate(ctx, input)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Empty(t, token)
	})
}
