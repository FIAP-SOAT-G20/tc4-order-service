package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
	mockport "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/usecase"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type CustomerUsecaseSuiteTest struct {
	suite.Suite
	// conn    *sql.DB
	// mock    sqlmock.Sqlmock
	// handler handler
	mockCustomers []*entity.Customer
	mockGateway   *mockport.MockCustomerGateway
	useCase       port.CustomerUseCase
	ctx           context.Context
}

func (s *CustomerUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockCustomerGateway(ctrl)
	s.useCase = usecase.NewCustomerUseCase(s.mockGateway)
	s.ctx = context.Background()
	currentTime := time.Now()
	s.mockCustomers = []*entity.Customer{
		{
			ID:        1,
			Name:      "Test Customer 1",
			Email:     "test.customer.1@email.com",
			CPF:       "12345678901",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
		{
			ID:        2,
			Name:      "Test Customer 2",
			Email:     "test.customer.2@email.com",
			CPF:       "12345678902",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}
}

// func (ts *CustomerUsecaseSuiteTest) AfterTest(_, _ string) {
// 	// assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
// }

func TestCustomerUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(CustomerUsecaseSuiteTest))
}
