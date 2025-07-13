package usecase_test

import (
	"testing"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *PaymentUsecaseSuiteTest) Test_paymentUseCase_Create() {
	tests := []struct {
		name        string
		input       dto.CreatePaymentInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Payment, error)
	}{
		{
			name:  "should create payment successfully",
			input: dto.CreatePaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().FindByOrderIDAndStatusProcessing(s.ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				s.mockOrderUseCase.EXPECT().Get(s.ctx, gomock.Any()).Return(&entity.Order{ID: uint64(1), OrderProducts: []entity.OrderProduct{{OrderID: 1, ProductID: 1}}}, nil)
				s.mockGateway.EXPECT().CreateExternal(s.ctx, gomock.Any()).Return(&entity.CreatePaymentExternalOutput{}, nil)
				s.mockGateway.EXPECT().Create(s.ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				s.mockOrderUseCase.EXPECT().Update(s.ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
			},
		},
		{
			name:  "should return error when update from order use case fails",
			input: dto.CreatePaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().FindByOrderIDAndStatusProcessing(s.ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				s.mockOrderUseCase.EXPECT().Get(s.ctx, gomock.Any()).Return(&entity.Order{ID: uint64(1), OrderProducts: []entity.OrderProduct{{OrderID: 1, ProductID: 1}}}, nil)
				s.mockGateway.EXPECT().CreateExternal(s.ctx, gomock.Any()).Return(&entity.CreatePaymentExternalOutput{}, nil)
				s.mockGateway.EXPECT().Create(s.ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				s.mockOrderUseCase.EXPECT().Update(s.ctx, gomock.Any()).Return(nil, &domain.InternalError{})
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name:  "should return error when create from gateway fails",
			input: dto.CreatePaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().FindByOrderIDAndStatusProcessing(s.ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				s.mockOrderUseCase.EXPECT().Get(s.ctx, gomock.Any()).Return(&entity.Order{ID: uint64(1), OrderProducts: []entity.OrderProduct{{OrderID: 1, ProductID: 1}}}, nil)
				s.mockGateway.EXPECT().CreateExternal(s.ctx, gomock.Any()).Return(&entity.CreatePaymentExternalOutput{}, nil)
				s.mockGateway.EXPECT().Create(s.ctx, gomock.Any()).Return(&entity.Payment{}, &domain.InternalError{})
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name:  "should return error when CreateExternal from gateway fails",
			input: dto.CreatePaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().FindByOrderIDAndStatusProcessing(s.ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				s.mockOrderUseCase.EXPECT().Get(s.ctx, gomock.Any()).Return(&entity.Order{ID: uint64(1), OrderProducts: []entity.OrderProduct{{OrderID: 1, ProductID: 1}}}, nil)
				s.mockGateway.EXPECT().CreateExternal(s.ctx, gomock.Any()).Return(&entity.CreatePaymentExternalOutput{}, assert.AnError)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name:  "should return error when dont have order product",
			input: dto.CreatePaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().FindByOrderIDAndStatusProcessing(s.ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				s.mockOrderUseCase.EXPECT().Get(s.ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, &domain.NotFoundError{}, err)
			},
		},
		{
			name:  "should return error when get from order use case fails",
			input: dto.CreatePaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().FindByOrderIDAndStatusProcessing(s.ctx, gomock.Any()).Return(&entity.Payment{}, nil)
				s.mockOrderUseCase.EXPECT().Get(s.ctx, gomock.Any()).Return(&entity.Order{}, assert.AnError)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, &domain.NotFoundError{}, err)
			},
		},
		{
			name:  "should return error when FindByOrderIDAndStatusProcessing from gateway fails",
			input: dto.CreatePaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().FindByOrderIDAndStatusProcessing(s.ctx, gomock.Any()).Return(&entity.Payment{}, assert.AnError)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name:  "should return the existing payment when already exists",
			input: dto.CreatePaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().FindByOrderIDAndStatusProcessing(s.ctx, gomock.Any()).Return(&entity.Payment{ID: 1}, nil)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
				assert.Equal(t, uint64(1), payment.ID)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			payment, err := s.useCase.Create(s.ctx, tt.input)

			tt.checkResult(t, payment, err)
		})
	}
}

func (s *PaymentUsecaseSuiteTest) Test_paymentUseCase_Get() {
	tests := []struct {
		name        string
		input       dto.GetPaymentInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Payment, error)
	}{
		{
			name:  "should return the payment successfully",
			input: dto.GetPaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByOrderID(s.ctx, gomock.Any()).
					Return(&entity.Payment{ID: 1}, nil)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
			},
		},
		{
			name:  "should return error when gateway fails",
			input: dto.GetPaymentInput{OrderID: uint64(1)},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByOrderID(s.ctx, gomock.Any()).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			payment, err := s.useCase.Get(s.ctx, tt.input)

			tt.checkResult(t, payment, err)
		})
	}
}

func (s *PaymentUsecaseSuiteTest) Test_paymentUseCase_Update() {
	tests := []struct {
		name        string
		input       dto.UpdatePaymentInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Payment, error)
	}{
		{
			name: "should update the payment",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().Update(s.ctx, gomock.Any(), gomock.Any()).Return(nil)
				s.mockGateway.EXPECT().FindByExternalPaymentID(s.ctx, gomock.Any()).Return(&entity.Payment{ID: 1}, nil)
				s.mockOrderUseCase.EXPECT().Get(s.ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
				s.mockOrderUseCase.EXPECT().Update(s.ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
			},
		},
		{
			name: "should return error when update from order use case fails",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().Update(s.ctx, gomock.Any(), gomock.Any()).Return(nil)
				s.mockGateway.EXPECT().FindByExternalPaymentID(s.ctx, gomock.Any()).Return(&entity.Payment{ID: 1}, nil)
				s.mockOrderUseCase.EXPECT().Get(s.ctx, gomock.Any()).Return(&entity.Order{ID: 1}, nil)
				s.mockOrderUseCase.EXPECT().Update(s.ctx, gomock.Any()).Return(nil, &domain.InternalError{})
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name: "should return error when get from order use case fails",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().Update(s.ctx, gomock.Any(), gomock.Any()).Return(nil)
				s.mockGateway.EXPECT().FindByExternalPaymentID(s.ctx, gomock.Any()).Return(&entity.Payment{ID: 1}, nil)
				s.mockOrderUseCase.EXPECT().Get(s.ctx, gomock.Any()).Return(&entity.Order{}, &domain.InternalError{})
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name: "should return error when FindByExternalPaymentID from gateway fails",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().Update(s.ctx, gomock.Any(), gomock.Any()).Return(nil)
				s.mockGateway.EXPECT().FindByExternalPaymentID(s.ctx, gomock.Any()).Return(&entity.Payment{}, assert.AnError)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, assert.AnError, err)
			},
		},
		{
			name: "should return error when gateway fails",
			input: dto.UpdatePaymentInput{
				Resource: "389d873a-436b-4ef2-a47a-0abf9b3e9924",
				Topic:    "payment",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().Update(s.ctx, gomock.Any(), gomock.Any()).Return(assert.AnError)
			},
			checkResult: func(t *testing.T, payment *entity.Payment, err error) {
				assert.Error(t, err)
				assert.Nil(t, payment)
				assert.IsType(t, assert.AnError, err)
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			payment, err := s.useCase.Update(s.ctx, tt.input)

			tt.checkResult(t, payment, err)
		})
	}
}
