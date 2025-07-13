package handler_test

import (
	"context"
	"testing"

	mockport "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/infrastructure/handler"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type StaffHandlerSuiteTest struct {
	suite.Suite
	handler        *handler.StaffHandler
	router         *gin.Engine
	mockController *mockport.MockStaffController
	ctx            context.Context
	requests       map[string]string // Fixture files
	responses      map[string]string // Golden files
}

func (s *StaffHandlerSuiteTest) SetupTest() {
	// Create a new router
	s.router = newRouter()

	// Create a new handler
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockController = mockport.NewMockStaffController(ctrl)
	s.handler = handler.NewStaffHandler(s.mockController)
	s.ctx = context.Background()

	// Register routes
	s.router.GET("/staffs", s.handler.List)
	s.router.POST("/staffs", s.handler.Create)
	s.router.PUT("/staffs/:id", s.handler.Update)
	s.router.GET("/staffs/:id", s.handler.Get)
	s.router.DELETE("/staffs/:id", s.handler.Delete)

	// Mock requests
	var err error
	s.requests, err = util.ReadFixtureFiles("staff",
		"create_success", "create_invalid_body",
		"update_success", "update_invalid_body",
	)
	assert.NoError(s.T(), err)

	// Mock responses
	s.responses, err = util.ReadGoldenFiles("staff",
		"list_success", "list_success_with_query",
		"create_success",
		"update_success",
		"get_success",
		"delete_success",
	)
	assert.NoError(s.T(), err)
	addCommonResponses(&s.responses)
}

func TestStaffHandlerSuiteTest(t *testing.T) {
	suite.Run(t, new(StaffHandlerSuiteTest))
}
