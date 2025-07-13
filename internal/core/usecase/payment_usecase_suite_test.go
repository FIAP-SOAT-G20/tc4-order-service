package usecase_test

import (
	"context"
	"testing"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
	mockport "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/usecase"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type PaymentUsecaseSuiteTest struct {
	suite.Suite
	mockGateway      *mockport.MockPaymentGateway
	mockOrderUseCase *mockport.MockOrderUseCase
	useCase          port.PaymentUseCase
	ctx              context.Context
}

func (s *PaymentUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockPaymentGateway(ctrl)
	s.mockOrderUseCase = mockport.NewMockOrderUseCase(ctrl)
	s.useCase = usecase.NewPaymentUseCase(s.mockGateway, s.mockOrderUseCase)
	s.ctx = context.Background()
}

func TestPaymentUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(PaymentUsecaseSuiteTest))
}
