package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
)

func (s *CustomerUsecaseSuiteTest) TestCustomersUseCase_List() {
	tests := []struct {
		name        string
		input       dto.ListCustomersInput
		setupMocks  func()
		checkResult func(*testing.T, []*entity.Customer, int64, error)
	}{
		{
			name: "should list products successfully",
			input: dto.ListCustomersInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", 1, 10).
					Return(s.mockCustomers, int64(2), nil)
			},
			checkResult: func(t *testing.T, customers []*entity.Customer, total int64, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, customers)
				assert.Equal(t, len(s.mockCustomers), len(customers))
				assert.Equal(t, int64(2), total)
			},
		},
		{
			name: "should return error when repository fails",
			input: dto.ListCustomersInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			checkResult: func(t *testing.T, customers []*entity.Customer, total int64, err error) {
				assert.Error(t, err)
				assert.Nil(t, customers)
				assert.Equal(t, int64(0), total)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name: "should filter by name",
			input: dto.ListCustomersInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "Test", 1, 10).
					Return(s.mockCustomers, int64(2), nil)
			},
			checkResult: func(t *testing.T, customers []*entity.Customer, total int64, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, customers)
				assert.Equal(t, len(s.mockCustomers), len(customers))
				assert.Equal(t, int64(2), total)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			customers, total, err := s.useCase.List(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, customers, total, err)
		})
	}
}

func (s *CustomerUsecaseSuiteTest) TestCustomerUseCase_Create() {
	tests := []struct {
		name        string
		input       dto.CreateCustomerInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Customer, error)
	}{
		{
			name: "should create customer successfully",
			input: dto.CreateCustomerInput{
				Name:  s.mockCustomers[0].Name,
				Email: s.mockCustomers[0].Email,
				CPF:   s.mockCustomers[0].CPF,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					Create(s.ctx, gomock.Any()).
					Return(nil)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, s.mockCustomers[0].Name, customer.Name)
				assert.Equal(t, s.mockCustomers[0].Email, customer.Email)
				assert.Equal(t, s.mockCustomers[0].CPF, customer.CPF)
			},
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreateCustomerInput{
				Name:  s.mockCustomers[0].Name,
				Email: s.mockCustomers[0].Email,
				CPF:   s.mockCustomers[0].CPF,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					Create(s.ctx, gomock.Any()).
					Return(assert.AnError)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.Error(t, err)
				assert.Nil(t, customer)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			customer, err := s.useCase.Create(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, customer, err)
		})
	}
}

func (s *CustomerUsecaseSuiteTest) TestCustomerUseCase_Get() {
	tests := []struct {
		name        string
		input       dto.GetCustomerInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Customer, error)
	}{
		{
			name:  "should get customer successfully",
			input: dto.GetCustomerInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(s.mockCustomers[0], nil)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, s.mockCustomers[0].ID, customer.ID)
				assert.Equal(t, s.mockCustomers[0].Name, customer.Name)
				assert.Equal(t, s.mockCustomers[0].Email, customer.Email)
				assert.Equal(t, s.mockCustomers[0].CreatedAt, customer.CreatedAt)
				assert.Equal(t, s.mockCustomers[0].UpdatedAt, customer.UpdatedAt)
			},
		},
		{
			name:  "should return not found error when customer doesn't exist",
			input: dto.GetCustomerInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.Error(t, err)
				assert.Nil(t, customer)
				assert.IsType(t, &domain.NotFoundError{}, err)
			},
		},
		{
			name:  "should return internal error when gateway fails",
			input: dto.GetCustomerInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.Error(t, err)
				assert.Nil(t, customer)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			customer, err := s.useCase.Get(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, customer, err)
		})
	}
}

func (s *CustomerUsecaseSuiteTest) TestCustomerUseCase_Update() {
	tests := []struct {
		name        string
		input       dto.UpdateCustomerInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Customer, error)
	}{
		{
			name: "should update customer successfully",
			input: dto.UpdateCustomerInput{
				ID:    1,
				Name:  "New Name",
				Email: "new.name@email.com",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(s.mockCustomers[0], nil)

				s.mockGateway.EXPECT().
					Update(s.ctx, gomock.Any()).
					DoAndReturn(func(_ context.Context, p *entity.Customer) error {
						assert.Equal(s.T(), "New Name", p.Name)
						assert.Equal(s.T(), "new.name@email.com", p.Email)
						return nil
					})
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, "New Name", customer.Name)
				assert.Equal(t, "new.name@email.com", customer.Email)
				assert.Equal(t, s.mockCustomers[0].CreatedAt, customer.CreatedAt)
			},
		},
		{
			name: "should return error when customer not found",
			input: dto.UpdateCustomerInput{
				ID:    1,
				Name:  "New Name",
				Email: "new.name@email.com",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.Error(t, err)
				assert.Nil(t, customer)
				assert.IsType(t, &domain.NotFoundError{}, err)
			},
		},
		{
			name: "should return error when gateway find fails",
			input: dto.UpdateCustomerInput{
				ID:    1,
				Name:  "New Name",
				Email: "new.name@email.com",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.Error(t, err)
				assert.Nil(t, customer)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name: "should return error when gateway update fails",
			input: dto.UpdateCustomerInput{
				ID:    1,
				Name:  "New Name",
				Email: "new.name@email.com",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(s.mockCustomers[0], nil)

				s.mockGateway.EXPECT().
					Update(s.ctx, gomock.Any()).
					Return(assert.AnError)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.Error(t, err)
				assert.Nil(t, customer)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			customer, err := s.useCase.Update(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, customer, err)
		})
	}
}

func (s *CustomerUsecaseSuiteTest) TestCustomerUseCase_Delete() {
	tests := []struct {
		name        string
		input       dto.DeleteCustomerInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Customer, error)
	}{
		{
			name:  "should delete customer successfully",
			input: dto.DeleteCustomerInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(&entity.Customer{ID: 1}, nil)

				s.mockGateway.EXPECT().
					Delete(s.ctx, uint64(1)).
					Return(nil)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, uint64(1), customer.ID)
			},
		},
		{
			name:  "should return not found error when customer doesn't exist",
			input: dto.DeleteCustomerInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.Error(t, err)
				assert.Nil(t, customer)
				assert.IsType(t, &domain.NotFoundError{}, err)
			},
		},
		{
			name:  "should return error when gateway fails on find",
			input: dto.DeleteCustomerInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.Error(t, err)
				assert.Nil(t, customer)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name:  "should return error when gateway fails on delete",
			input: dto.DeleteCustomerInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(&entity.Customer{}, nil)

				s.mockGateway.EXPECT().
					Delete(s.ctx, uint64(1)).
					Return(assert.AnError)
			},
			checkResult: func(t *testing.T, customer *entity.Customer, err error) {
				assert.Error(t, err)
				assert.Nil(t, customer)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			customer, err := s.useCase.Delete(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, customer, err)
		})
	}
}
