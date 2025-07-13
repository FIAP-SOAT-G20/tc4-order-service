package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/util"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *StaffHandlerSuiteTest) TestStaffHandler_List() {
	tests := []struct {
		name        string
		url         string
		setupMocks  func()
		checkResult func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success",
			url:  "/staffs",
			setupMocks: func() {
				s.mockController.EXPECT().List(gomock.Any(), gomock.Any(), dto.ListStaffsInput{
					Page:  1,
					Limit: 10,
				}).Return([]byte(s.responses["list_success"]), nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Contains(t, res.Body.String(), s.responses["list_success"])
			},
		},
		{
			name:       "invalid query - page",
			url:        "/staffs?page=invalid",
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_invalid_query"])
			},
		},
		{
			name: "success - with query - role",
			url:  "/staffs?role=COOK",
			setupMocks: func() {
				s.mockController.EXPECT().List(gomock.Any(), gomock.Any(), dto.ListStaffsInput{
					Role:  valueobject.COOK,
					Page:  1,
					Limit: 10,
				}).Return([]byte(s.responses["list_success_with_query"]), nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Contains(t, res.Body.String(), s.responses["list_success_with_query"])
			},
		},
		{
			name: "controller error",
			url:  "/staffs",
			setupMocks: func() {
				s.mockController.EXPECT().List(gomock.Any(), gomock.Any(), dto.ListStaffsInput{
					Page:  1,
					Limit: 10,
				}).Return(nil, domain.NewInternalError(nil))
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_internal_error"])
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)

			// Act
			s.router.ServeHTTP(w, req)

			// Assert
			tt.checkResult(t, w)
		})
	}
}

func (s *StaffHandlerSuiteTest) TestStaffHandler_Create() {
	tests := []struct {
		name        string
		url         string
		body        *strings.Reader
		setupMocks  func()
		checkResult func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success",
			url:  "/staffs",
			body: strings.NewReader(s.requests["create_success"]),
			setupMocks: func() {
				s.mockController.EXPECT().
					Create(gomock.Any(), gomock.Any(), dto.CreateStaffInput{
						Name: "John Doe 6",
						Role: valueobject.COOK,
					}).
					Return([]byte(s.responses["create_success"]), nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["create_success"])
			},
		},
		{
			name:       "invalid request - body is not a valid json",
			url:        "/staffs",
			body:       strings.NewReader("invalid"),
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
			},
		},
		{
			name:       "invalid request - body filed Name is a number",
			url:        "/staffs",
			body:       strings.NewReader(s.requests["create_invalid_body"]),
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
			},
		},
		{
			name: "controller error",
			url:  "/staffs",
			body: strings.NewReader(s.requests["create_success"]),
			setupMocks: func() {
				s.mockController.EXPECT().
					Create(gomock.Any(), gomock.Any(), dto.CreateStaffInput{
						Name: "John Doe 6",
						Role: valueobject.COOK,
					}).
					Return(nil, domain.NewInternalError(nil))
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_internal_error"])
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tt.url, tt.body)

			// Act
			s.router.ServeHTTP(w, req)

			// Assert
			tt.checkResult(t, w)
		})
	}
}

func (s *StaffHandlerSuiteTest) TestStaffHandler_Get() {
	tests := []struct {
		name        string
		url         string
		setupMocks  func()
		checkResult func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success",
			url:  "/staffs/6",
			setupMocks: func() {
				s.mockController.EXPECT().
					Get(gomock.Any(), gomock.Any(), dto.GetStaffInput{ID: 6}).
					Return([]byte(s.responses["get_success"]), nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["get_success"])
			},
		},
		{
			name: "not found",
			url:  "/staffs/5",
			setupMocks: func() {
				s.mockController.EXPECT().
					Get(gomock.Any(), gomock.Any(), dto.GetStaffInput{ID: 5}).
					Return(nil, domain.NewNotFoundError(domain.ErrNotFound))
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_not_found"])
			},
		},
		{
			name:       "invalid request - id is not a number",
			url:        "/staffs/invalid",
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_invalid_param"])
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)

			// Act
			s.router.ServeHTTP(w, req)

			// Assert
			tt.checkResult(t, w)
		})
	}
}

func (s *StaffHandlerSuiteTest) TestStaffHandler_Update() {
	tests := []struct {
		name        string
		url         string
		body        *strings.Reader
		setupMocks  func()
		checkResult func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success - update Staff status",
			url:  "/staffs/6",
			body: strings.NewReader(s.requests["update_success"]),
			setupMocks: func() {
				s.mockController.EXPECT().
					Update(gomock.Any(), gomock.Any(), dto.UpdateStaffInput{
						ID:   6,
						Name: "John Doe 6 UPDATED",
						Role: valueobject.ATTENDANT,
					}).
					Return([]byte(s.responses["update_success"]), nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["update_success"])
			},
		},
		{
			name:       "invalid request - body is not a valid json",
			url:        "/staffs/5",
			body:       strings.NewReader("invalid"),
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_invalid_body"])
			},
		},
		{
			name:       "invalid request - body field Name is a number",
			url:        "/staffs/5",
			body:       strings.NewReader(s.requests["update_invalid_body"]),
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_invalid_body"])
			},
		},
		{
			name:       "invalid request - id is not a number",
			url:        "/staffs/invalid",
			body:       strings.NewReader(s.requests["update_success"]),
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_invalid_param"])
			},
		},
		{
			name: "controller error",
			url:  "/staffs/6",
			body: strings.NewReader(s.requests["update_success"]),
			setupMocks: func() {
				s.mockController.EXPECT().
					Update(gomock.Any(), gomock.Any(), dto.UpdateStaffInput{
						ID:   6,
						Name: "John Doe 6 UPDATED",
						Role: valueobject.ATTENDANT,
					}).
					Return(nil, domain.NewInternalError(nil))
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_internal_error"])
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPut, tt.url, tt.body)

			// Act
			s.router.ServeHTTP(w, req)

			// Assert
			tt.checkResult(t, w)
		})
	}
}

func (s *StaffHandlerSuiteTest) TestStaffHandler_Delete() {
	tests := []struct {
		name        string
		url         string
		setupMocks  func()
		checkResult func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success",
			url:  "/staffs/6",
			setupMocks: func() {
				s.mockController.EXPECT().
					Delete(gomock.Any(), gomock.Any(), dto.DeleteStaffInput{ID: 6}).
					Return([]byte(s.responses["delete_success"]), nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["delete_success"])
			},
		},
		{
			name: "not found",
			url:  "/staffs/6",
			setupMocks: func() {
				s.mockController.EXPECT().
					Delete(gomock.Any(), gomock.Any(), dto.DeleteStaffInput{ID: 6}).
					Return(nil, domain.NewNotFoundError(domain.ErrNotFound))
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_not_found"])
			},
		},
		{
			name:       "invalid request - id is not a number",
			url:        "/staffs/invalid",
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Contains(t, util.RemoveAllSpaces(res.Body.String()), s.responses["error_invalid_param"])
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, tt.url, nil)

			// Act
			s.router.ServeHTTP(w, req)

			// Assert
			tt.checkResult(t, w)
		})
	}
}
