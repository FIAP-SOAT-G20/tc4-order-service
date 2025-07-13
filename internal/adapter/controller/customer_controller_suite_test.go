package controller_test

import (
	"context"
	"testing"
	"time"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
	mockport "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type CustomerControllerSuiteTest struct {
	suite.Suite
	mockUseCase   *mockport.MockCustomerUseCase
	mockPresenter port.Presenter
	mockCustomer  *entity.Customer
	mockCustomers []*entity.Customer
	controller    port.CustomerController
	ctx           context.Context
	// responses      map[string]string // Golden files
}

func (s *CustomerControllerSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockUseCase = mockport.NewMockCustomerUseCase(ctrl)
	s.mockPresenter = presenter.NewCustomerJsonPresenter()
	mockDateAt2, _ := time.Parse(time.RFC3339, "2025-03-06T17:03:28Z")
	mockDateAt3, _ := time.Parse(time.RFC3339, "2025-03-06T17:03:58Z")
	s.mockCustomer = &entity.Customer{
		ID:        6,
		Name:      "John Doe 6",
		Email:     "john.doe.6@email.com",
		CPF:       "000.000.000-06",
		CreatedAt: mockDateAt2,
		UpdatedAt: mockDateAt3,
	}
	mockDateAt, _ := time.Parse(time.RFC3339, "2025-02-28T16:28:18Z")
	s.mockCustomers = []*entity.Customer{
		{
			ID:        1,
			Name:      "John Doe 1",
			Email:     "john.doe.1@email.com",
			CPF:       "123.456.789-00",
			CreatedAt: mockDateAt,
			UpdatedAt: mockDateAt,
		},
		{
			ID:        2,
			Name:      "John Doe 2",
			Email:     "john.doe.2@email.com",
			CPF:       "987.654.321-00",
			CreatedAt: mockDateAt,
			UpdatedAt: mockDateAt,
		},
	}
	s.controller = controller.NewCustomerController(s.mockUseCase)
	s.ctx = context.Background()
}

func TestCustomerControllerSuiteTest(t *testing.T) {
	suite.Run(t, new(CustomerControllerSuiteTest))
}
