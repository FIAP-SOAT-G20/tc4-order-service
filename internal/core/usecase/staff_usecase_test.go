package usecase_test

import (
	"testing"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *StaffUsecaseSuiteTest) TestStaffsUseCase_List() {
	tests := []struct {
		name        string
		input       dto.ListStaffsInput
		setupMocks  func()
		checkResult func(*testing.T, []*entity.Staff, int64, error)
	}{
		{
			name: "should list staffs successfully",
			input: dto.ListStaffsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				var role valueobject.StaffRole
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", role, 1, 10).
					Return(s.mockStaffs, int64(2), nil)
			},
			checkResult: func(t *testing.T, staffs []*entity.Staff, total int64, err error) {
				assert.NoError(t, err)
				assert.Equal(t, s.mockStaffs, staffs)
				assert.Equal(t, int64(2), total)
			},
		},
		{
			name: "should return error when repository fails",
			input: dto.ListStaffsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				var role valueobject.StaffRole
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", role, 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			checkResult: func(t *testing.T, staffs []*entity.Staff, total int64, err error) {
				assert.Error(t, err)
				assert.Nil(t, staffs)
				assert.Equal(t, int64(0), total)
			},
		},
		{
			name: "should filter by name",
			input: dto.ListStaffsInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				var role valueobject.StaffRole
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "Test", role, 1, 10).
					Return(s.mockStaffs, int64(2), nil)
			},
			checkResult: func(t *testing.T, staffs []*entity.Staff, total int64, err error) {
				assert.NoError(t, err)
				assert.Equal(t, s.mockStaffs, staffs)
				assert.Equal(t, int64(2), total)
			},
		},
		{
			name: "should filter by Role",
			input: dto.ListStaffsInput{
				Role:  "COOK",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", valueobject.COOK, 1, 10).
					Return(s.mockStaffs, int64(2), nil)

			},
			checkResult: func(t *testing.T, staffs []*entity.Staff, total int64, err error) {
				assert.NoError(t, err)
				assert.Equal(t, s.mockStaffs, staffs)
				assert.Equal(t, int64(2), total)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			staffs, total, err := s.useCase.List(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, staffs, total, err)
		})
	}
}

func (s *StaffUsecaseSuiteTest) TestStaffUseCase_Create() {
	tests := []struct {
		name        string
		input       dto.CreateStaffInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Staff, error)
	}{
		{
			name: "should create staff successfully",
			input: dto.CreateStaffInput{
				Name: "John Smith",
				Role: "COOK",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					Create(s.ctx, gomock.Any()).
					Return(nil)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, staff)
				assert.Equal(t, "John Smith", staff.Name)
				assert.Equal(t, "COOK", string(staff.Role))
			},
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreateStaffInput{
				Name: "John Smith",
				Role: "COOK",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					Create(s.ctx, gomock.Any()).
					Return(assert.AnError)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.Error(t, err)
				assert.Nil(t, staff)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			staff, err := s.useCase.Create(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, staff, err)
		})
	}
}

func (s *StaffUsecaseSuiteTest) TestStaffUseCase_Get() {
	tests := []struct {
		name        string
		input       dto.GetStaffInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Staff, error)
	}{
		{
			name:  "should get staff successfully",
			input: dto.GetStaffInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(s.mockStaffs[0], nil)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, staff)
				assert.Equal(t, uint64(1), staff.ID)
			},
		},
		{
			name:  "should return not found error when staff doesn't exist",
			input: dto.GetStaffInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.Error(t, err)
				assert.Nil(t, staff)
				assert.IsType(t, &domain.NotFoundError{}, err)
			},
		},
		{
			name:  "should return internal error when gateway fails",
			input: dto.GetStaffInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.Error(t, err)
				assert.Nil(t, staff)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			staff, err := s.useCase.Get(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, staff, err)
		})
	}
}

func (s *StaffUsecaseSuiteTest) TestStaffUseCase_Update() {
	tests := []struct {
		name        string
		input       dto.UpdateStaffInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Staff, error)
	}{
		{
			name: "should update staff successfully",
			input: dto.UpdateStaffInput{
				ID:   1,
				Name: "New Name",
				Role: "ATTENDANT",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(s.mockStaffs[0], nil)

				s.mockGateway.EXPECT().
					Update(s.ctx, gomock.Any()).
					Return(nil)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, staff)
				assert.Equal(t, uint64(1), staff.ID)
				assert.Equal(t, "New Name", staff.Name)
				assert.Equal(t, "ATTENDANT", string(staff.Role))
			},
		},
		{
			name: "should return error when staff not found",
			input: dto.UpdateStaffInput{
				ID:   1,
				Name: "New Name",
				Role: "ATTENDANT",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.Error(t, err)
				assert.Nil(t, staff)
				assert.IsType(t, &domain.NotFoundError{}, err)
			},
		},
		{
			name: "should return error when gateway find fails",
			input: dto.UpdateStaffInput{
				ID:   1,
				Name: "New Name",
				Role: "ATTENDANT",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.Error(t, err)
				assert.Nil(t, staff)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name: "should return error when gateway update fails",
			input: dto.UpdateStaffInput{
				ID:   1,
				Name: "New Name",
				Role: "ATTENDANT",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(s.mockStaffs[0], nil)

				s.mockGateway.EXPECT().
					Update(s.ctx, gomock.Any()).
					Return(assert.AnError)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.Error(t, err)
				assert.Nil(t, staff)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			staff, err := s.useCase.Update(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, staff, err)
		})
	}
}

func (s *StaffUsecaseSuiteTest) TestStaffUseCase_Delete() {
	tests := []struct {
		name        string
		input       dto.DeleteStaffInput
		setupMocks  func()
		checkResult func(*testing.T, *entity.Staff, error)
	}{
		{
			name:  "should delete staff successfully",
			input: dto.DeleteStaffInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(s.mockStaffs[0], nil)

				s.mockGateway.EXPECT().
					Delete(s.ctx, uint64(1)).
					Return(nil)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, staff)
				assert.Equal(t, uint64(1), staff.ID)
			},
		},
		{
			name:  "should return not found error when staff doesn't exist",
			input: dto.DeleteStaffInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.Error(t, err)
				assert.Nil(t, staff)
				assert.IsType(t, &domain.NotFoundError{}, err)
			},
		},
		{
			name:  "should return error when gateway fails on find",
			input: dto.DeleteStaffInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.Error(t, err)
				assert.Nil(t, staff)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
		{
			name:  "should return error when gateway fails on delete",
			input: dto.DeleteStaffInput{ID: 1},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(&entity.Staff{}, nil)

				s.mockGateway.EXPECT().
					Delete(s.ctx, uint64(1)).
					Return(assert.AnError)
			},
			checkResult: func(t *testing.T, staff *entity.Staff, err error) {
				assert.Error(t, err)
				assert.Nil(t, staff)
				assert.IsType(t, &domain.InternalError{}, err)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()

			// Act
			staff, err := s.useCase.Delete(s.ctx, tt.input)

			// Assert
			tt.checkResult(t, staff, err)
		})
	}
}
