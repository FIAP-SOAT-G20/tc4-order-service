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

type StaffUsecaseSuiteTest struct {
	suite.Suite
	mockStaffs  []*entity.Staff
	mockGateway *mockport.MockStaffGateway
	useCase     port.StaffUseCase
	ctx         context.Context
}

func (s *StaffUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockStaffGateway(ctrl)
	s.useCase = usecase.NewStaffUseCase(s.mockGateway)
	s.ctx = context.Background()
	currentTime := time.Now()
	s.mockStaffs = []*entity.Staff{
		{
			ID:        1,
			Name:      "Test Staff 1",
			Role:      "COOK",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
		{
			ID:        2,
			Name:      "Test Staff 2",
			Role:      "COOK",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}
}

func TestStaffUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(StaffUsecaseSuiteTest))
}
