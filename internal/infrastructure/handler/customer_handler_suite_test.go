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

type CustomerHandlerSuiteTest struct {
	suite.Suite
	handler        *handler.CustomerHandler
	router         *gin.Engine
	mockController *mockport.MockCustomerController
	ctx            context.Context
	requests       map[string]string // Fixture files
	responses      map[string]string // Golden files
}

func (s *CustomerHandlerSuiteTest) SetupTest() {
	// Create a new router
	s.router = newRouter()

	// Create a new handler
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockController = mockport.NewMockCustomerController(ctrl)
	s.handler = handler.NewCustomerHandler(s.mockController)
	s.ctx = context.Background()

	// Register routes
	s.router.GET("/customers", s.handler.List)
	s.router.POST("/customers", s.handler.Create)
	s.router.PUT("/customers/:id", s.handler.Update)
	s.router.GET("/customers/:id", s.handler.Get)
	s.router.DELETE("/customers/:id", s.handler.Delete)

	// Mock requests
	var err error
	s.requests, err = util.ReadFixtureFiles("customer",
		"create_success", "create_invalid_body",
		"update_success", "update_invalid_body",
	)
	assert.NoError(s.T(), err)

	// Mock responses
	s.responses, err = util.ReadGoldenFiles("customer",
		"list_success", "list_success_with_query",
		"create_success",
		"update_success",
		"get_success",
		"delete_success",
	)
	assert.NoError(s.T(), err)
	addCommonResponses(&s.responses)
}

func TestCustomerHandlerSuiteTest(t *testing.T) {
	suite.Run(t, new(CustomerHandlerSuiteTest))
}
