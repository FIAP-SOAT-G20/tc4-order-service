package controller_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/util"
)

func (s *CustomerControllerSuiteTest) TestCustomerController_ListCustomers() {
	tests := []struct {
		name        string
		input       dto.ListCustomersInput
		setupMocks  func()
		checkResult func(*testing.T, []byte, error)
	}{
		{
			name: "List customers success",
			input: dto.ListCustomersInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					List(s.ctx, dto.ListCustomersInput{
						Name:  "Test",
						Page:  1,
						Limit: 10,
					}).
					Return(s.mockCustomers, int64(2), nil)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				want, _ := util.ReadGoldenFile("customer/list_success")
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, want, util.RemoveAllSpaces(string(output)))
			},
		},
		{
			name: "List customers use case error",
			input: dto.ListCustomersInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					List(s.ctx, dto.ListCustomersInput{
						Name:  "Test",
						Page:  1,
						Limit: 10,
					}).
					Return(nil, int64(0), assert.AnError)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				assert.Error(t, err)
				assert.Nil(t, output)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			output, err := s.controller.List(s.ctx, s.mockPresenter, tt.input)

			// Assert
			tt.checkResult(t, output, err)
		})
	}
}

func (s *CustomerControllerSuiteTest) TestCustomerController_CreateCustomer() {
	tests := []struct {
		name        string
		input       dto.CreateCustomerInput
		setupMocks  func()
		checkResult func(*testing.T, []byte, error)
	}{
		{
			name: "Create customers success",
			input: dto.CreateCustomerInput{
				Name:  s.mockCustomer.Name,
				Email: s.mockCustomer.Email,
				CPF:   s.mockCustomer.CPF,
			},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					Create(s.ctx, dto.CreateCustomerInput{
						Name:  s.mockCustomer.Name,
						Email: s.mockCustomer.Email,
						CPF:   s.mockCustomer.CPF,
					}).
					Return(s.mockCustomer, nil)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				want, _ := util.ReadGoldenFile("customer/create_success")
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, want, util.RemoveAllSpaces(string(output)))
			},
		},
		{
			name: "Create use case error",
			input: dto.CreateCustomerInput{
				Name:  s.mockCustomer.Name,
				Email: s.mockCustomer.Email,
				CPF:   s.mockCustomer.CPF,
			},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					Create(s.ctx, dto.CreateCustomerInput{
						Name:  s.mockCustomer.Name,
						Email: s.mockCustomer.Email,
						CPF:   s.mockCustomer.CPF,
					}).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				assert.Error(t, err)
				assert.Nil(t, output)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			output, err := s.controller.Create(s.ctx, s.mockPresenter, tt.input)

			// Assert
			tt.checkResult(t, output, err)
		})
	}
}

func (s *CustomerControllerSuiteTest) TestCustomerController_GetCustomer() {
	tests := []struct {
		name        string
		input       dto.GetCustomerInput
		setupMocks  func()
		checkResult func(*testing.T, []byte, error)
	}{
		{
			name:  "Get customers success",
			input: dto.GetCustomerInput{ID: uint64(1)},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					Get(s.ctx, dto.GetCustomerInput{
						ID: uint64(1),
					}).
					Return(s.mockCustomer, nil)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				want, _ := util.ReadGoldenFile("customer/get_success")
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, want, util.RemoveAllSpaces(string(output)))
			},
		},
		{
			name:  "Get use case error",
			input: dto.GetCustomerInput{ID: uint64(1)},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					Get(s.ctx, dto.GetCustomerInput{ID: uint64(1)}).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				assert.Error(t, err)
				assert.Nil(t, output)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			output, err := s.controller.Get(s.ctx, s.mockPresenter, tt.input)

			// Assert
			tt.checkResult(t, output, err)
		})
	}
}

func (s *CustomerControllerSuiteTest) TestCustomerController_UpdateCustomer() {
	customerUpdated := &entity.Customer{
		ID:        6,
		Name:      "John Doe 6 UPDATED",
		Email:     "john.doe.6.updated@email.com",
		CPF:       s.mockCustomer.CPF,
		CreatedAt: s.mockCustomer.CreatedAt,
		UpdatedAt: s.mockCustomer.UpdatedAt,
	}
	tests := []struct {
		name        string
		input       dto.UpdateCustomerInput
		setupMocks  func()
		checkResult func(*testing.T, []byte, error)
	}{
		{
			name: "Update customers success",
			input: dto.UpdateCustomerInput{
				ID:    customerUpdated.ID,
				Name:  customerUpdated.Name,
				Email: customerUpdated.Email,
			},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					Update(s.ctx, dto.UpdateCustomerInput{
						ID:    customerUpdated.ID,
						Name:  customerUpdated.Name,
						Email: customerUpdated.Email,
					}).
					Return(customerUpdated, nil)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				want, _ := util.ReadGoldenFile("customer/update_success")
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, want, util.RemoveAllSpaces(string(output)))
			},
		},
		// {
		// 	name: "Update use case error",
		// },
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			output, err := s.controller.Update(s.ctx, s.mockPresenter, tt.input)

			// Assert
			tt.checkResult(t, output, err)
		})
	}
}

func (s *CustomerControllerSuiteTest) TestCustomerController_DeleteCustomer() {
	tests := []struct {
		name        string
		input       dto.DeleteCustomerInput
		setupMocks  func()
		checkResult func(*testing.T, []byte, error)
	}{
		{
			name:  "Delete customers success",
			input: dto.DeleteCustomerInput{ID: uint64(1)},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					Delete(s.ctx, dto.DeleteCustomerInput{ID: uint64(1)}).
					Return(s.mockCustomer, nil)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				want, _ := util.ReadGoldenFile("customer/delete_success")
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, want, util.RemoveAllSpaces(string(output)))
			},
		},
		{
			name:  "Delete use case error",
			input: dto.DeleteCustomerInput{ID: uint64(1)},
			setupMocks: func() {
				s.mockUseCase.EXPECT().
					Delete(s.ctx, dto.DeleteCustomerInput{ID: uint64(1)}).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, output []byte, err error) {
				assert.Error(t, err)
				assert.Nil(t, output)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			output, err := s.controller.Delete(s.ctx, s.mockPresenter, tt.input)

			// Assert
			tt.checkResult(t, output, err)
		})
	}
}
