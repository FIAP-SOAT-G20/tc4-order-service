package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/infrastructure/handler/request"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/infrastructure/middleware"
)

type PaymentHandler struct {
	controller port.PaymentController
	jwtService port.JWTService
}

func NewPaymentHandler(controller port.PaymentController, jwtService port.JWTService) *PaymentHandler {
	return &PaymentHandler{controller, jwtService}
}

func (h *PaymentHandler) Register(router *gin.RouterGroup) {
	router.POST("/:order_id/checkout", h.Create, middleware.JWTAuthMiddleware(h.jwtService))
	router.POST("/callback", h.Update)
	router.GET("/:order_id", h.Get)
}

// Create godoc
//
//	@Summary		Create a payment (Checkout) (Reference TC-1 2.b.v; TC-2 1.a.i, 1.a.v)
//	@Description	Creates a new payment (Checkout)
//	@Description	The status of the payment will be set to PROCESSING
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			order_id						path		int								true	"Order ID"
//	@Success		201								{object}	presenter.PaymentJsonResponse	"Created"
//	@Failure		400								{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		500								{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/payments/{order_id}/checkout	[post]
func (h *PaymentHandler) Create(c *gin.Context) {
	var uri request.CreatePaymentUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.CreatePaymentInput{
		OrderID: uri.OrderID,
	}

	output, err := h.controller.Create(
		c.Request.Context(),
		presenter.NewPaymentJsonPresenter(),
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
//	@Summary		Update a payment (Webhook) (Reference TC-2 1.a.iii)
//	@Description	Update a new payment (Webhook)
//	@Description	- resource = external payment id, obtained from the checkout response
//	@Description	- topic = payment
//	@Description
//	@Description	> The status of the payment will be set to CONFIRMED if the payment was successful
//	@Description	## Possible status:
//	@Description	- `PROCESSING` (default)
//	@Description	- `CONFIRMED`
//	@Description	- `FAILED`
//	@Description	- `ABORTED`
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			payment				body		request.UpdatePaymentRequest	true	"Payment data"
//	@Success		201					{object}	presenter.PaymentJsonResponse	"Created"
//	@Failure		400					{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		500					{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/payments/callback	[post]
func (h *PaymentHandler) Update(c *gin.Context) {

	var body request.UpdatePaymentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdatePaymentInput{
		Resource: body.Resource,
		Topic:    body.Topic,
	}

	output, err := h.controller.Update(
		c.Request.Context(),
		presenter.NewPaymentJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}

// Get godoc
//
//	@Summary		Get a payment given order ID (Reference TC-2 1.a.ii)
//	@Description	Get a payment given order ID
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			order_id				path		int								true	"Order ID"
//	@Success		201						{object}	presenter.PaymentJsonResponse	"Created"
//	@Failure		400						{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		500						{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/payments/{order_id}	[get]
func (h *PaymentHandler) Get(c *gin.Context) {

	var body request.GetPaymentRequest
	if err := c.ShouldBindUri(&body); err != nil {
		fmt.Println(err)
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.GetPaymentInput{
		OrderID: body.OrderID,
	}

	output, err := h.controller.Get(
		c.Request.Context(),
		presenter.NewPaymentJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}
