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

type AuthHandler struct {
	controller port.AuthController
}

func NewAuthHandler(controller port.AuthController) *AuthHandler {
	return &AuthHandler{controller}
}

func (h *AuthHandler) Register(router *gin.RouterGroup) {
	router.POST("", h.Authenticate)
}

// Authenticate godoc
//
//	@Summary		Authenticate user
//	@Description	Authenticates a user by CPF and returns a JWT token
//	@Tags			sign-in
//	@Accept			json
//	@Produce		json
//	@Param			authentication	body		request.AuthenticateBodyRequest		true	"User CPF"
//	@Success		200				{object}	presenter.AuthenticationResponse	"OK"
//	@Failure		400				{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Failure		401				{object}	middleware.ErrorJsonResponse		"Unauthorized"
//	@Failure		404				{object}	middleware.ErrorJsonResponse		"Not Found"
//	@Failure		500				{object}	middleware.ErrorJsonResponse		"Internal Server Error"
//	@Router			/auth [post]
func (h *AuthHandler) Authenticate(c *gin.Context) {
	var body request.AuthenticateBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.AuthenticateInput{
		CPF: body.CPF,
	}

	output, err := h.controller.Authenticate(
		c.Request.Context(),
		presenter.NewAuthPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}
