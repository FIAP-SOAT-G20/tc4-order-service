package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/infrastructure/handler/request"
)

type StaffHandler struct {
	controller port.StaffController
}

func NewStaffHandler(controller port.StaffController) *StaffHandler {
	return &StaffHandler{controller: controller}
}

func (h *StaffHandler) Register(router *gin.RouterGroup) {
	router.GET("", h.List)
	router.POST("", h.Create)
	router.GET("/:id", h.Get)
	router.PUT("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
}

// List godoc
//
//	@Summary		List staffs
//	@Description	List all staffs
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string									false	"Filter by name"
//	@Param			role	query		string									false	"Filter by role. Available options: COOK, ATTENDANT, MANAGER"
//	@Param			page	query		int										false	"Page number"		default(1)
//	@Param			limit	query		int										false	"Items per page"	default(10)
//	@Success		200		{object}	presenter.StaffJsonPaginatedResponse	"OK"
//	@Failure		400		{object}	middleware.ErrorJsonResponse			"Bad Request"
//	@Failure		500		{object}	middleware.ErrorJsonResponse			"Internal Server Error"
//	@Router			/staffs [get]
func (h *StaffHandler) List(c *gin.Context) {
	var query request.ListStaffsQueryRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.ListStaffsInput{
		Name:  query.Name,
		Role:  query.Role,
		Page:  query.Page,
		Limit: query.Limit,
	}

	output, err := h.controller.List(
		c.Request.Context(),
		presenter.NewStaffJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}

// Create godoc
//
//	@Summary		Create staff
//	@Description	Creates a new staff
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			staff	body		request.CreateStaffBodyRequest	true	"Staff data"
//	@Success		201		{object}	presenter.StaffJsonResponse		"Created"
//	@Failure		400		{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		500		{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/staffs [post]
func (h *StaffHandler) Create(c *gin.Context) {
	var body request.CreateStaffBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.CreateStaffInput{
		Name: body.Name,
		Role: body.Role,
	}

	output, err := h.controller.Create(
		c.Request.Context(),
		presenter.NewStaffJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusCreated, "application/json", output)
}

// Get godoc
//
//	@Summary		Get staff
//	@Description	Search for a staff by ID
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"Staff ID"
//	@Success		200	{object}	presenter.StaffJsonResponse		"OK"
//	@Failure		400	{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/staffs/{id} [get]
func (h *StaffHandler) Get(c *gin.Context) {
	var uri request.GetStaffUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetStaffInput{
		ID: uri.ID,
	}

	output, err := h.controller.Get(
		c.Request.Context(),
		presenter.NewStaffJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}

// Update godoc
//
//	@Summary		Update staff
//	@Description	Update an existing staff
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"Staff ID"
//	@Param			staff	body		request.UpdateStaffBodyRequest	true	"Staff data"
//	@Success		200		{object}	presenter.StaffJsonResponse		"OK"
//	@Failure		400		{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404		{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500		{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/staffs/{id} [put]
func (h *StaffHandler) Update(c *gin.Context) {
	var uri request.UpdateStaffUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var body request.UpdateStaffBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdateStaffInput{
		ID:   uri.ID,
		Name: body.Name,
		Role: body.Role,
	}

	output, err := h.controller.Update(
		c.Request.Context(),
		presenter.NewStaffJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}

// Delete godoc
//
//	@Summary		Delete staff
//	@Description	Deletes a staff by ID
//	@Tags			staffs
//	@Produce		json
//	@Param			id	path		int								true	"Staff ID"
//	@Success		200	{object}	presenter.StaffJsonResponse		"OK"
//	@Failure		400	{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/staffs/{id} [delete]
func (h *StaffHandler) Delete(c *gin.Context) {
	var uri request.DeleteStaffUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteStaffInput{
		ID: uri.ID,
	}
	output, err := h.controller.Delete(
		c.Request.Context(),
		presenter.NewStaffJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}
