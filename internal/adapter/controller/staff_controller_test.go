package controller_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	mockport "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port/mocks"
)

// TODO: Add more test cenarios
func TestStaffController_ListStaffs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStaffsUseCase := mockport.NewMockStaffUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := controller.NewStaffController(mockStaffsUseCase)

	ctx := context.Background()
	input := dto.ListStaffsInput{
		Name:  "Test",
		Role:  "COOK",
		Page:  1,
		Limit: 10,
	}

	currentTime := time.Now()
	mockStaffs := []*entity.Staff{
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

	mockStaffsUseCase.EXPECT().
		List(ctx, input).
		Return(mockStaffs, int64(2), nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{
			Result: mockStaffs,
			Total:  int64(2),
			Page:   1,
			Limit:  10,
		}).
		Return([]byte{}, nil)

	output, err := controller.List(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestStaffController_CreateStaff(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStaffUseCase := mockport.NewMockStaffUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := controller.NewStaffController(mockStaffUseCase)

	ctx := context.Background()
	input := dto.CreateStaffInput{
		Name: "Test Staff",
		Role: "COOK",
	}

	mockStaff := &entity.Staff{
		ID:   1,
		Name: "Test Staff",
		Role: "COOK",
	}

	mockStaffUseCase.EXPECT().
		Create(ctx, input).
		Return(mockStaff, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockStaff}).
		Return([]byte{}, nil)

	output, err := controller.Create(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestStaffController_GetStaff(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStaffUseCase := mockport.NewMockStaffUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := controller.NewStaffController(mockStaffUseCase)

	ctx := context.Background()
	input := dto.GetStaffInput{
		ID: uint64(1),
	}

	mockStaff := &entity.Staff{
		ID:   1,
		Name: "Test Staff",
		Role: "COOK",
	}

	mockStaffUseCase.EXPECT().
		Get(ctx, input).
		Return(mockStaff, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockStaff}).
		Return([]byte{}, nil)

	output, err := controller.Get(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestStaffController_UpdateStaff(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStaffUseCase := mockport.NewMockStaffUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := controller.NewStaffController(mockStaffUseCase)

	ctx := context.Background()
	input := dto.UpdateStaffInput{
		ID:   uint64(1),
		Name: "Staff UPDATED",
		Role: "ATTENDANT",
	}

	mockStaff := &entity.Staff{
		ID:   1,
		Name: "Updated Staff",
		Role: "ATTENDANT",
	}

	mockStaffUseCase.EXPECT().
		Update(ctx, input).
		Return(mockStaff, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockStaff}).
		Return([]byte{}, nil)

	output, err := controller.Update(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestStaffController_DeleteStaff(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStaffUseCase := mockport.NewMockStaffUseCase(ctrl)
	mockPresenter := mockport.NewMockPresenter(ctrl)
	controller := controller.NewStaffController(mockStaffUseCase)

	ctx := context.Background()
	input := dto.DeleteStaffInput{
		ID: uint64(1),
	}

	mockStaff := &entity.Staff{
		ID:   1,
		Name: "Test Staff",
		Role: "COOK",
	}

	mockStaffUseCase.EXPECT().
		Delete(ctx, input).
		Return(mockStaff, nil)

	mockPresenter.EXPECT().
		Present(dto.PresenterInput{Result: mockStaff}).
		Return([]byte{}, nil)

	output, err := controller.Delete(ctx, mockPresenter, input)
	assert.NoError(t, err)
	assert.NotNil(t, output)
}
